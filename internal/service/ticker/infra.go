package ticker

import (
	"context"
	"fmt"
	"log/slog"
	"rss/internal/domain"
	"rss/internal/ports/inbound"
	"rss/internal/ports/outbound"
	"rss/internal/service/worker"
	"time"
)

type ticker struct {
	ctx      context.Context
	interval time.Duration
	jobs     chan *domain.FeedForGetReq
	db       outbound.PsqlForTicker
	workers  worker.WorkerResizeForTicker
}

func BuildTicker(db outbound.PsqlForTicker, workers worker.WorkerResizeForTicker) inbound.TickController {
	return &ticker{
		db:      db,
		workers: workers,
	}
}

func (t *ticker) Start(ctx context.Context, inter time.Duration) (<-chan *domain.FeedForGetReq, error) {
	if t.ctx != nil {
		return nil, fmt.Errorf("")
	}
	signal, err := t.db.Listen(ctx)
	if err != nil {
		return nil, err
	}
	t.ctx = ctx
	t.interval = inter
	t.jobs = make(chan *domain.FeedForGetReq, 64)
	go t.startTick(signal)
	return t.jobs, nil
}

func (t *ticker) startTick(signal <-chan struct{}) {
	tick := time.NewTicker(t.interval)
	defer tick.Stop()
	defer close(t.jobs)
	for {
		select {
		case <-t.ctx.Done():
			return
		case <-signal:
			stg, err := t.db.GetSettings(t.ctx)
			if err != nil {
				slog.Error("ticker", "get setting", err)
				return
			}
			if !stg.IsRunning {
				t.workers.StopAll()
				return
			} else {
				go t.workers.ResizeWorker(stg.Worker)
			}
			if t.interval != stg.Interval {
				tick.Reset(stg.Interval)
			}
		case <-tick.C:
			feeds, err := t.db.ListRssWithLastUpdatedChan(t.ctx)
			if err != nil {
				slog.Error("ticker", "get listRSS", err)
				return
			}
			for feed := range feeds {
				t.jobs <- feed
			}
		}
	}
}
