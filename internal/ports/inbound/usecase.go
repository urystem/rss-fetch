package inbound

import (
	"context"
	"rss/internal/domain"
	"time"
)

type UseCasePsql interface {
	Starter(ctx context.Context) error
	Stopper(ctx context.Context) error

	ResizeWorker(ctx context.Context, workers uint) (uint, error)
	SetInterval(ctx context.Context, d time.Duration) (time.Duration, error)

	RssAdd(ctx context.Context, name, url string) error
	ListRssFeeds(ctx context.Context) ([]domain.Feed, error)
	ListRssFeedsWithNum(ctx context.Context, n uint) ([]domain.Feed, error)
	DeleteRssFeed(ctx context.Context, name string) error

	ShowArticles(ctx context.Context, name string, n uint) ([]domain.Article, error)
}
