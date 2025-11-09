package inbound

import (
	"context"
	"rss/internal/domain"
)

type Workers interface {
	StopAll()
	WorkerForCLI
	ResizeWorker(count int)
}

type WorkerForCLI interface {
	Start(ctx context.Context, workers int, jobs <-chan *domain.FeedForGetReq) error
}
