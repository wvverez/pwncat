package config

import (
	"encoding/json"
	"os"
	"time"
)

const Version = "1.0.1"

type Config struct {
	// Objetivo
	URL      string
	Wordlist string
	Method   string

	// Rendimiento
	Threads int
	Rate    int
	Timeout time.Duration
	Retry   int
	Delay   time.Duration

	// Filtros (match)
	MatchStatus string
	MatchSize   string
	MatchRegex  string

	// Filtros (exclude)
	ExcludeStatus string
	ExcludeSize   string
	ExcludeRegex  string

	// Salida
	Output   string
	Format   string
	NoColor  bool
	Verbose  bool
	DebugLog string

	// Avanzados
	ConfigFile string
	Replay     string
	Cert       string
	Key        string
	Insecure   bool
	ShowVersion bool
}

func New() *Config {
	return &Config{
		Method:  "GET",
		Threads: 40,
		Timeout: 10 * time.Second,
		Retry:   3,
		Format:  "json",
	}
}

func LoadFile(cfg *Config) error {
	data, err := os.ReadFile(cfg.ConfigFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, cfg)
}
