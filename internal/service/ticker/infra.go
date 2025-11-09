package ticker

import (
	"context"
	"log/slog"
	"rss/internal/domain"
	"rss/internal/ports/inbound"
	"rss/internal/ports/outbound"
	"rss/internal/service/worker"
	"time"
)

type ticker struct {
	ctx     context.Context
	cancel  context.CancelFunc
	logger  *slog.Logger
	jobs    chan *domain.FeedForGetReq
	db      outbound.PsqlForTicker
	workers worker.WorkerResizeForTicker
}

func BuildTicker(ctx context.Context, cancel context.CancelFunc, logger *slog.Logger, db outbound.PsqlForTicker, workers worker.WorkerResizeForTicker) inbound.TickController {
	return &ticker{
		ctx:     ctx,
		cancel:  cancel,
		logger:  logger,
		db:      db,
		workers: workers,
	}
}

func (t *ticker) Start(inter time.Duration) (<-chan *domain.FeedForGetReq, error) {
	signal, err := t.db.Listen(t.ctx)
	if err != nil {
		return nil, err
	}
	t.jobs = make(chan *domain.FeedForGetReq, 64)
	go t.startTick(signal, inter)
	return t.jobs, nil
}

func (t *ticker) Stop(ctx context.Context) {
	if t.jobs != nil {
		t.db.Stopper(ctx)
		t.logger.Info("running status = false")
	}
}
