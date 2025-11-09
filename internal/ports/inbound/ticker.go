package inbound

import (
	"context"
	"rss/internal/domain"
	"time"
)

type TickController interface {
	Start(inter time.Duration) (<-chan *domain.FeedForGetReq, error)
	Stop(ctx context.Context)
}
