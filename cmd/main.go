package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"rss/internal/bootstrap"
	"rss/internal/configs"
	"syscall"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	ctxBack := context.Background()
	cfg := configs.Load()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app, err := bootstrap.InitApp(ctxBack, cfg, logger)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Run(ctxBack); err != nil {
			slog.Error("❌", " Server error:", err)
			quit <- syscall.SIGTERM
		} else {
			close(quit)
		}
	}()
	// time.Sleep(20 * time.Second)
	// close(quit)
	_, ok := <-quit // Ждём сигнал
	if !ok {
		return
	}
	slog.Info("📦 Shutting down server...")

	if err := app.Shutdown(ctxBack); err != nil {
		slog.Error("❌", " Server forced to shutdown: %v", err)
	}
	slog.Info("✅ Server exited properly")
}

func usage() {
	fmt.Println(`
Usage:

    rsshub COMMAND [OPTIONS]

Common Commands:

    add             Add new RSS feed
    set-interval    Set RSS fetch interval
    set-workers     Set number of workers
    list            List available RSS feeds
    delete          Delete RSS feed
    articles        Show latest articles
    fetch           Start background process that periodically fetches and processes RSS feeds using a worker pool`)
}
