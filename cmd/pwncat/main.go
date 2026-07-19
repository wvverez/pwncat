package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"pwncat/internal/config"
	"pwncat/internal/engine"
	"pwncat/internal/filter"
	"pwncat/internal/http"
	"pwncat/internal/input"
	"pwncat/internal/output"
)

func main() {
	cfg := config.New()
	cfg.ParseFlags()

	if cfg.ShowVersion {
		fmt.Printf("pwncat version: %s\n", config.Version)
		os.Exit(0)
	}

	if cfg.DebugLog != "" {
		f, err := os.OpenFile(cfg.DebugLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("Warning: Could not open debug log: %s", err)
		} else {
			log.SetOutput(f)
			defer f.Close()
		}
	}

	if cfg.ConfigFile != "" {
		if err := config.LoadFile(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %s\n", err)
			os.Exit(1)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	source, err := input.NewProvider(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	matcher, err := filter.NewMatcher(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	excluder, err := filter.NewExcluder(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	client := http.NewClient(cfg)
	display := output.NewDisplay(cfg)

	eng := engine.New(cfg, source, matcher, excluder, client, display)

	if err := eng.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
