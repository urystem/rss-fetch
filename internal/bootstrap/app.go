package bootstrap

import (
	"context"
	"log/slog"
	"rss/internal/adapters/driven/storage"
	"rss/internal/configs"
	"rss/internal/ports/inbound"
	"rss/internal/ports/outbound"
)

type myApp struct {
	db     outbound.PostgresInter
	logger *slog.Logger
	wcfg   configs.WorkerInter
}

func InitApp(ctx context.Context, cfg configs.ConfigInter, logger *slog.Logger) (inbound.AppInter, error) {
	myDB, err := storage.InitDB(ctx, cfg.GetDBConfig())
	if err != nil {
		return nil, err
	}
	return &myApp{
		db:     myDB,
		logger: logger,
		wcfg:   cfg.GetWorkerCfg(),
	}, nil
}

func (app *myApp) Run(ctx context.Context) error {
	return nil
}

func (app *myApp) Shutdown(ctx context.Context) error {
	return nil
}
