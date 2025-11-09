package inbound

import (
	"context"
	"rss/internal/domain"
	"time"
)

type TickController interface {
	Start(ctx context.Context, cancelMain context.CancelCauseFunc, inter time.Duration) (<-chan *domain.FeedForGetReq, error)
}
