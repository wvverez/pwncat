package engine

import (
	"context"
	"sync"
	"time"

	"pwncat/internal/config"
	"pwncat/internal/filter"
	"pwncat/internal/http"
	"pwncat/internal/input"
	"pwncat/internal/output"
)

type Engine struct {
	cfg       *config.Config
	source    input.Provider
	matcher   filter.Condition
	excluder  filter.Condition
	client    *http.Client
	display   *output.Display
	stats     *output.Stats
	results   chan output.Result
	statsMu   sync.Mutex
	mu        sync.Mutex
	paused    bool
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
}

func New(cfg *config.Config, source input.Provider, matcher, excluder filter.Condition, client *http.Client, display *output.Display) *Engine {
	return &Engine{
		cfg:      cfg,
		source:   source,
		matcher:  matcher,
		excluder: excluder,
		client:   client,
		display:  display,
		stats:    &output.Stats{StartTime: time.Now()},
		results:  make(chan output.Result, 1000),
	}
}

func (e *Engine) Run(ctx context.Context) error {
	e.ctx, e.cancel = context.WithCancel(ctx)
	defer e.cancel()

	e.display.Start()
	defer e.display.Stop()

	var wg sync.WaitGroup
	done := make(chan bool)

	go e.processResults(&wg, done)

	for i := 0; i < e.cfg.Threads; i++ {
		wg.Add(1)
		go e.worker(&wg)
	}

	wg.Add(1)
	go e.monitorProgress(&wg)

	wg.Wait()
	close(e.results)
	<-done

	e.stats.EndTime = time.Now()
	e.display.FinalStats(e.stats)

	if e.cfg.Output != "" {
		e.saveResults()
	}

	return nil
}






func (e *Engine) worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var limiter *time.Ticker
	if e.cfg.Rate > 0 {
		limiter = time.NewTicker(time.Second / time.Duration(e.cfg.Rate))
		defer limiter.Stop()
	}

	for {
		select {
		case <-e.ctx.Done():
			return
		default:
			if e.paused {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			if limiter != nil {
				select {
				case <-limiter.C:
				case <-e.ctx.Done():
					return
				}
			}

			val, err := e.source.Next()
			if err != nil {
				return
			}

			e.processInput(val)
		}
	}
}






func (e *Engine) processInput(val *input.Value) {
	req := http.NewRequest(e.cfg)
	req.SetValue(val)

	resp, err := e.client.Do(req)
	if err != nil {
		e.statsMu.Lock()
		e.stats.Errors++
		e.statsMu.Unlock()
		e.results <- output.Result{Input: string(val.Values["PWN"]), Error: err}
		return
	}

	e.statsMu.Lock()
	e.stats.Total++
	e.statsMu.Unlock()

	matched := e.matcher.Match(resp)
	filtered := e.excluder.Match(resp)

	if matched && !filtered {
		e.statsMu.Lock()
		e.stats.Matched++
		e.statsMu.Unlock()
		e.results <- output.Result{Input: string(val.Values["PWN"]), Response: resp, Matched: true}
	} else {
		e.statsMu.Lock()
		e.stats.Filtered++
		e.statsMu.Unlock()
		e.results <- output.Result{Input: string(val.Values["PWN"]), Response: resp, Matched: false}
	}
}

func (e *Engine) processResults(wg *sync.WaitGroup, done chan bool) {
	defer func() { done <- true }()
	wg.Add(1)
	defer wg.Done()

	for result := range e.results {
		select {
		case <-e.ctx.Done():
			return
		default:
			if result.Error != nil {
				e.display.Error(result)
			} else {
				e.display.Result(result)
			}
		}
	}
}

func (e *Engine) monitorProgress(wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-e.ctx.Done():
			return
		case <-ticker.C:
			e.display.Progress(e.stats)
		}
	}
}

func (e *Engine) Pause() {
	e.mu.Lock()
	e.paused = true
	e.mu.Unlock()
}

func (e *Engine) Resume() {
	e.mu.Lock()
	e.paused = false
	e.mu.Unlock()
}

func (e *Engine) saveResults() {

}
