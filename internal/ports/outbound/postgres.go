package outbound

import (
	"context"
	"rss/internal/domain"
	"time"
)

type PostgresInter interface {
	PsqlSetting
	PsqlRssFeed
	PsqlArticles
}

type PsqlSetting interface {
	SetInterval(ctx context.Context, d time.Duration) (time.Duration, error)
	ResizeWorker(ctx context.Context, workers uint) (uint, error)
	Starter(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type PsqlRssFeed interface {
	RssAdd(ctx context.Context, name, url string) error
	ListRssFeeds(ctx context.Context) ([]domain.Feed, error)
	ListRssFeedsWithNum(ctx context.Context, n uint) ([]domain.Feed, error)
	DeleteRssFeed(ctx context.Context, name string) error

	ListRssWithLastUpdated(ctx context.Context, n uint) ([]domain.Feed, error)
	UpdateFeedsLastUpdates(ctx context.Context, feeds []string) error
}

type PsqlArticles interface {
	ShowArticles(ctx context.Context, name string, n uint) ([]domain.Article, error)
	InsertArticles(ctx context.Context, feedID string, articles []domain.Article) error
}
