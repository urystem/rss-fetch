package outbound

import (
	"context"
	"rss/internal/domain"
)

type RssHttp interface {
	GetRss(ctx context.Context, url string) ([]domain.RSSItem, error)
}
