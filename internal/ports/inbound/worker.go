package inbound

import (
	"context"
	"rss/internal/domain"
)

type Workers interface {
	Start(ctx context.Context, workers int, jobs <-chan *domain.FeedForGetReq) error
	StopAll()
	WorkerResize
}

type WorkerResize interface {
	ResizeWorker(count int)
}
