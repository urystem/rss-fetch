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

func InitApp(ctxMain context.Context, cfg configs.ConfigInter, logger *slog.Logger) (inbound.AppInter, error) {
	ctx, cancel := context.WithCancel(ctxMain)
	myDB, err := storage.InitDB(ctx, cfg.GetDBConfig())
	if err != nil {
		cancel()
		return nil, err
	}
	rss := rss.BuildRss(5 * time.Second)
	work := worker.BuildWorker(logger, myDB, rss)
	tick := ticker.BuildTicker(ctx, cancel, logger, myDB, work)
	use := usecase.BuildBridge(myDB, tick, work, cfg.GetWorkerCfg())
	cli := cli.BuildCli(use)
	return &myApp{
		ctx:     ctx,
		cancel:  cancel,
		db:      myDB,
		logger:  logger,
		workers: work,
		tick:    tick,
		cli:     cli,
		cfg:     cfg.GetWorkerCfg(),
	}, nil
}

func (app *myApp) Run() error {
	// defer app.Shutdown(ctx)
	return app.cli.Run(app.ctx)
}

func (app *myApp) Shutdown(ctx context.Context) error {
	app.cancel()
	app.tick.Stop(ctx)
	app.db.CloseDB()
	fmt.Println("db closed")
	return nil
}
