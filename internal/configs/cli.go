package configs

import (
	"log/slog"
	"os"
	"strconv"
	"time"
)

type worker struct {
	interval    *time.Duration
	countWorker *uint
}

type WorkerInter interface {
	GetWorkerCount() *uint
	GetInterval() *time.Duration
}

func initWorkerCfg() worker {
	var (
		in    *time.Duration
		count *uint
	)
	interStr := os.Getenv("CLI_APP_TIMER_INTERVAL")
	if interStr != "" {
		ti, err := time.ParseDuration(interStr)
		if err != nil {
			slog.Error("config", "worker", "invalid interval")
			os.Exit(1)
		}
		in = &ti
	}
	countStr := os.Getenv("CLI_APP_WORKERS_COUNT")
	if countStr != "" {
		co, err := strconv.ParseUint(countStr, 10, 64) // 10 — основание, 64 — битность
		if err != nil {
			slog.Error("config", "worker", "invalid count worker")
			os.Exit(1)
		}
		temp := uint(co)
		count = &temp
	}
	return worker{
		interval:    in,
		countWorker: count,
	}
}

func (w *worker) GetWorkerCount() *uint { return w.countWorker }

func (w *worker) GetInterval() *time.Duration { return w.interval }
