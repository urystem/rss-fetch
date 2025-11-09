package worker

import (
	"context"
	"fmt"
	"log/slog"
	"rss/internal/domain"
	"rss/internal/ports/inbound"
	"rss/internal/ports/outbound"
	"sync"
)

type workersDo struct {
	ctx        context.Context
	controller []chan struct{}
	workerMu   sync.RWMutex
	logger     slog.Logger
	db         outbound.PsqlForWorkers
	rss        outbound.RssHttp
	jobs       <-chan *domain.FeedForGetReq
}

func BuildWorker(logg slog.Logger, db outbound.PsqlForWorkers, rss outbound.RssHttp) inbound.Workers {
	return &workersDo{
		logger: logg,
		db:     db,
		rss:    rss,
	}
}

func (w *workersDo) Start(ctx context.Context, workers int, jobs <-chan *domain.FeedForGetReq) error {
	if w.ctx != nil {
		return fmt.Errorf("already running")
	}
	w.ctx = ctx
	w.jobs = jobs
	go w.stopCtx()
	w.ResizeWorker(workers)
	return nil
}

func (w *workersDo) ResizeWorker(count int) {
	if w.ctx == nil || w.ctx.Err() != nil {
		return
	}
	w.workerMu.Lock()
	defer w.workerMu.Unlock()
	diff := count - len(w.controller)
	if diff > 0 { //add
		for range diff {
			w.addWorker()
		}
	} else {
		for range -diff {
			w.delWorker()
		}
	}
}

func (w *workersDo) StopAll() {
	w.workerMu.Lock()
	defer w.workerMu.Unlock()
	for range len(w.controller) {
		w.delWorker()
	}
}
