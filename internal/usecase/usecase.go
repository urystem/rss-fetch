package usecase

import (
	"context"
	"fmt"
	"rss/internal/configs"
	"rss/internal/domain"
	"rss/internal/ports/inbound"
	"rss/internal/ports/outbound"
	"time"
)

type psqlUseCase struct {
	db   outbound.PsqlForCli
	tick inbound.TickController
	work inbound.WorkerForCLI
	cfg  configs.WorkerInter
}

func BuildBridge(db outbound.PsqlForCli, tick inbound.TickController, work inbound.WorkerForCLI, cfg configs.WorkerInter) inbound.UseCasePsql {
	return &psqlUseCase{
		db:   db,
		tick: tick,
		work: work,
		cfg:  cfg,
	}
}

func (p *psqlUseCase) Starter(ctxMain context.Context) error {
	ctx, cancel := context.WithCancelCause(ctxMain)
	defer cancel(fmt.Errorf("error build"))
	defer p.Stopper(ctx)

	err := p.db.Starter(ctx)
	if err != nil {
		return err
	}
	countWorker := p.cfg.GetWorkerCount()
	interval := p.cfg.GetInterval()
	err = p.db.SetAndGetSettings(ctx, &countWorker, &interval)
	if err != nil {
		return err
	}
	jobs, err := p.tick.Start(ctx, cancel, *interval)
	if err != nil {
		return err
	}
	err = p.work.Start(ctx, int(*countWorker), jobs)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return ctx.Err()
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
