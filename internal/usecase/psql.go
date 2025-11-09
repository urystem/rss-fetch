package usecase

import (
	"context"
	"rss/internal/domain"
	"rss/internal/ports/inbound"
	"rss/internal/ports/outbound"
	"time"
)

type psqlUseCase struct {
	db outbound.PsqlForCli
}

func BuildBridge(db outbound.PsqlForCli) inbound.UseCasePsql {
	return &psqlUseCase{
		db: db,
	}
}

func (p *psqlUseCase) Starter(ctx context.Context) error {
	return p.db.Starter(ctx)
}

func (p *psqlUseCase) SetAndGetSettings(ctx context.Context, workerCount *uint, intrv *time.Duration) error {
	return p.db.SetAndGetSettings(ctx, workerCount, intrv)
}

func (p *psqlUseCase) ResizeWorker(ctx context.Context, workers uint) (uint, error) {
	return p.db.ResizeWorker(ctx, workers)
}

func (p *psqlUseCase) SetInterval(ctx context.Context, d time.Duration) (time.Duration, error) {
	return p.db.SetInterval(ctx, d)
}

func (p *psqlUseCase) RssAdd(ctx context.Context, name, url string) error {
	return p.db.RssAdd(ctx, name, url)
}

func (p *psqlUseCase) ListRssFeeds(ctx context.Context) ([]domain.Feed, error) {
	return p.db.ListRssFeeds(ctx)
}

func (p *psqlUseCase) ListRssFeedsWithNum(ctx context.Context, n uint) ([]domain.Feed, error) {
	return p.db.ListRssFeedsWithNum(ctx, n)
}

func (p *psqlUseCase) DeleteRssFeed(ctx context.Context, name string) error {
	return p.db.DeleteRssFeed(ctx, name)
}

func (p *psqlUseCase) ShowArticles(ctx context.Context, name string, n uint) ([]domain.Article, error) {
	return p.db.ShowArticles(ctx, name, n)
}

func (p *psqlUseCase) Stopper(ctx context.Context) error {
	return p.db.Stopper(ctx)
}
