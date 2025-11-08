package configs

import (
	"log/slog"
	"os"
	"time"
)

type worker struct {
	interval    time.Duration
	countWorker uint
}

type WorkerInter interface {
	GetWorkerCount() uint
	GetInterval() time.Duration
}

func initWorkerCfg() worker {
	ti, err := time.ParseDuration(mustGetEnvString("CLI_APP_TIMER_INTERVAL"))
	if err != nil {
		slog.Error("config", "worker", "invalid interval")
		os.Exit(1)
	}
	return worker{
		interval:    ti,
		countWorker: uint(mustGetEnvInt("CLI_APP_WORKERS_COUNT")),
	}
}

func (w *worker) GetWorkerCount() uint { return w.countWorker }

func (w *worker) GetInterval() time.Duration { return w.interval }
