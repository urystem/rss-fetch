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
	ctx        context.Context
	cancelMain context.CancelCauseFunc
	logger     *slog.Logger
	jobs       chan *domain.FeedForGetReq
	db         outbound.PsqlForTicker
	workers    worker.WorkerResizeForTicker
}

func BuildTicker(logger *slog.Logger, db outbound.PsqlForTicker, workers worker.WorkerResizeForTicker) inbound.TickController {
	return &ticker{
		logger:  logger,
		db:      db,
		workers: workers,
	}
}

func (t *ticker) Start(ctx context.Context, cancelMain context.CancelCauseFunc, inter time.Duration) (<-chan *domain.FeedForGetReq, error) {
	if t.ctx != nil {
		return nil, fmt.Errorf("")
	}
	signal, err := t.db.Listen(ctx)
	if err != nil {
		return nil, err
	}
	t.ctx = ctx
	t.cancelMain = cancelMain
	t.jobs = make(chan *domain.FeedForGetReq, 64)
	go t.startTick(signal, inter)
	return t.jobs, nil
}
