package engine

import (
	"context"
	"time"
)

type Worker struct {
	id     int
	engine *Engine
	ctx    context.Context
	cancel context.CancelFunc
}

func NewWorker(id int, engine *Engine) *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		id:     id,
		engine: engine,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (w *Worker) Start() {
	go w.run()
}

func (w *Worker) run() {
	for {
		select {
		case <-w.ctx.Done():
			return
		default:
			if w.engine.paused {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			val, err := w.engine.source.Next()
			if err != nil {
				return
			}

			w.engine.processInput(val)
		}
	}
}

func (w *Worker) Stop() {
	w.cancel()
}
