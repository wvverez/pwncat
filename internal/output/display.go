package output

import (
	"fmt"
	"strings"
	"time"

	"pwncat/internal/config"
	"pwncat/internal/http"
	"pwncat/pkg/utils"
)

const pwncatArt = `
                                       
▄▄▄▄  ▄▄   ▄▄ ▄▄  ▄▄  ▄▄▄▄  ▄▄▄ ▄▄▄▄▄▄ 
██▄█▀ ██ ▄ ██ ███▄██ ██▀▀▀ ██▀██  ██   
██     ▀█▀█▀  ██ ▀██ ▀████ ██▀██  ██   
`

type Result struct {
	Input    string
	Response *http.Response
	Matched  bool
	Error    error
}

type Stats struct {
	Total     int64
	Matched   int64
	Filtered  int64
	Errors    int64
	StartTime time.Time
	EndTime   time.Time
}

type Display struct {
	cfg *config.Config
}

func NewDisplay(cfg *config.Config) *Display {
	return &Display{cfg: cfg}
}

func (d *Display) Start() {
	fmt.Print(pwncatArt)
	fmt.Printf("\npwncat started at %s\n", time.Now().Format("15:04:05"))
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("URL         : %s\n", d.cfg.URL)
	fmt.Printf("Method      : %s\n", d.cfg.Method)
	fmt.Printf("Wordlist    : %s\n", d.cfg.Wordlist)
	fmt.Printf("Threads     : %d\n", d.cfg.Threads)
	fmt.Printf("Rate        : %d req/s\n", d.cfg.Rate)
	if d.cfg.MatchStatus != "" {
		fmt.Printf("Match Status: %s\n", d.cfg.MatchStatus)
	}
	if d.cfg.ExcludeStatus != "" {
		fmt.Printf("Exclude     : %s\n", d.cfg.ExcludeStatus)
	}
	if d.cfg.Output != "" {
		fmt.Printf("Output      : %s\n", d.cfg.Output)
	}
	fmt.Println(strings.Repeat("-", 100))
}

func (d *Display) Stop() {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("pwncat finished at %s\n", time.Now().Format("15:04:05"))
}

func (d *Display) Result(r Result) {
	code := r.Response.StatusCode
	color := d.getColor(code)
	status := fmt.Sprintf("%d", code)

	input := r.Input
	if len(input) > 30 {
		input = input[:27] + "..."
	}

	// Limpiar línea de progreso SIN salto de línea extra
	fmt.Printf("\r%s", strings.Repeat(" ", 80))

	if d.cfg.NoColor {
		fmt.Printf("\r🐱 PWN -> %-25s [Status: %s, Bytes: %d]\n", input, status, r.Response.Size)
	} else {
		fmt.Printf("\r🐱 PWN -> %-25s [Status: %s, Bytes: %d]\n", input, color(status), r.Response.Size)
	}
}

func (d *Display) Progress(s *Stats) {
	elapsed := time.Since(s.StartTime)
	rate := float64(s.Total) / elapsed.Seconds()
	fmt.Printf("\r[%d req] %d matches | %d errors | %.1f req/s    ",
		s.Total, s.Matched, s.Errors, rate)
}

func (d *Display) Error(r Result) {
	if d.cfg.Verbose {
		fmt.Printf("[ERR] %s: %v\n", r.Input, r.Error)
	}
}

func (d *Display) FinalStats(s *Stats) {
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("Total requests:  %d\n", s.Total)
	fmt.Printf("Matches:         %d\n", s.Matched)
	fmt.Printf("Filtered:        %d\n", s.Filtered)
	fmt.Printf("Errors:          %d\n", s.Errors)
	fmt.Printf("Time:            %s\n", s.EndTime.Sub(s.StartTime))
	if s.Total > 0 {
		rate := float64(s.Total) / s.EndTime.Sub(s.StartTime).Seconds()
		fmt.Printf("Requests/sec:    %.1f\n", rate)
	}
	fmt.Println(strings.Repeat("-", 100))
}

func (d *Display) getColor(code int) func(string) string {
	if d.cfg.NoColor {
		return func(s string) string { return s }
	}
	switch {
	case code >= 200 && code < 300:
		return utils.GreenText
	case code >= 300 && code < 400:
		return utils.YellowText
	case code >= 400 && code < 500:
		return utils.RedText
	default:
		return utils.MagentaText
	}
}
