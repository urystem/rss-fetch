package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"rss/internal/adapters/driven/rss"
	"rss/internal/adapters/driven/storage"
	"rss/internal/adapters/driver/cli"
	"rss/internal/configs"
	"rss/internal/ports/inbound"
	"rss/internal/ports/outbound"
	"rss/internal/service/ticker"
	"rss/internal/service/worker"
	"rss/internal/usecase"
	"time"
)

type myApp struct {
	ctx     context.Context
	cancel  context.CancelFunc
	db      outbound.PostgresInter
	logger  *slog.Logger
	workers inbound.Workers
	tick    inbound.TickController
	cli     inbound.CliInter
	cfg     configs.WorkerInter
}

func InitApp(ctx context.Context, cfg configs.ConfigInter, logger *slog.Logger) (inbound.AppInter, error) {
	myDB, err := storage.InitDB(ctx, cfg.GetDBConfig())
	if err != nil {
		return nil, err
	}
	rss := rss.BuildRss(5 * time.Second)
	work := worker.BuildWorker(logger, myDB, rss)
	tick := ticker.BuildTicker(logger, myDB, work)
	use := usecase.BuildBridge(myDB, tick, work, cfg.GetWorkerCfg())
	cli := cli.BuildCli(use)
	return &myApp{
		db:      myDB,
		logger:  logger,
		workers: work,
		tick:    tick,
		cli:     cli,
		cfg:     cfg.GetWorkerCfg(),
	}, nil
}

func (app *myApp) Run(ctx context.Context) error {
	defer app.Shutdown(ctx)
	if app.ctx != nil {
		return fmt.Errorf("already running")
	}
	app.ctx, app.cancel = context.WithCancel(ctx)

	return app.cli.Run(app.ctx)
}

func (app *myApp) Shutdown(ctx context.Context) error {
	app.cancel()
	time.Sleep(5 * time.Second)
	return nil
}
