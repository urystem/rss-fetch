package outbound

import (
	"context"
	"rss/internal/domain"
	"time"
)

type PostgresInter interface {
	PsqlForCli
	PsqlRssFeed
	PsqlArticles
	CloseDB()
}

type psqlSettingForCli interface {
	SetInterval(ctx context.Context, d time.Duration) (time.Duration, error)
	ResizeWorker(ctx context.Context, workers uint) (uint, error)
	Starter(ctx context.Context) error
	Stopper(ctx context.Context) error
	SetAndGetSettings(ctx context.Context, workerCount **uint, intrv **time.Duration) error
}

type PsqlForCli interface {
	psqlRssFeedUnion
	psqlgetArticles
	psqlSettingForCli
}

type psqlRssFeedUnion interface {
	RssAdd(ctx context.Context, name, url string) error
	ListRssFeeds(ctx context.Context) ([]domain.Feed, error)
	ListRssFeedsWithNum(ctx context.Context, n uint) ([]domain.Feed, error)
	DeleteRssFeed(ctx context.Context, name string) error
}

type PsqlRssFeed interface {
	psqlRssFeedUnion
	psqlRssFeedForWorkers
}

type PsqlForTicker interface {
	ListRssWithLastUpdatedChan(ctx context.Context) (<-chan *domain.FeedForGetReq, error)
	// ListRssWithLastUpdated(ctx context.Context, n uint) ([]domain.FeedForGetReq, error)
	Listen(ctx context.Context) (<-chan struct{}, error)
	GetSettings(ctx context.Context) (*domain.Setting, error)
	Stopper(ctx context.Context) error
}

type psqlRssFeedForWorkers interface {
	UpdateFeedsLastUpdates(ctx context.Context, feeds []string) error
}

type psqlgetArticles interface {
	ShowArticles(ctx context.Context, name string, n uint) ([]domain.Article, error)
}

type PsqlArticles interface {
	psqlgetArticles
	psqlArticlesForWorkers
	PsqlForTicker
}
type psqlArticlesForWorkers interface {
	InsertArticles(ctx context.Context, feedID string, articles []domain.RSSItem) error
}

type PsqlForWorkers interface {
	psqlRssFeedForWorkers
	psqlArticlesForWorkers
}
