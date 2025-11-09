package storage

import (
	"context"
	"rss/internal/domain"
	"time"
)

func (p *poolDB) Starter(ctx context.Context) error {
	var isRunnig bool
	const query = `
        UPDATE setting
        SET is_running = TRUE
        WHERE mine = TRUE
		RETURNING old.is_running`
	err := p.QueryRow(ctx, query).Scan(&isRunnig)
	if err != nil {
		return err
	}
	if isRunnig {
		return domain.ErrNotAffected
	}
	return nil
}

func (p *poolDB) SetAndGetSettings(ctx context.Context, workerCount *uint, intrv *time.Duration) error {
	var w uint
	var i time.Duration

	err := p.QueryRow(ctx, `
        UPDATE setting
        SET
            workers = COALESCE($1, workers),
            interval = COALESCE($2, interval)
        WHERE mine = TRUE AND is_running = TRUE
        RETURNING workers, interval
    `, workerCount, intrv).Scan(&w, &i)

	if err != nil {
		return err
	}

	// если нужно — обновляем входные параметры
	if workerCount != nil {
		*workerCount = w
	}
	if intrv != nil {
		*intrv = i
	}

	return nil
}

func (p *poolDB) GetSettings(ctx context.Context) (*domain.Setting, error) {
	const query = `
		SELECT is_running, workers, interval
		FROM setting
		WHERE mine = TRUE`

	s := new(domain.Setting)

	return s, p.QueryRow(ctx, query).Scan(
		s.IsRunning,
		s.Worker,
		s.Interval,
	)
}

func (p *poolDB) Shutdown(ctx context.Context) error {
	res, err := p.Exec(ctx, `
        UPDATE setting
        SET is_running = FALSE
        WHERE mine = TRUE`)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return domain.ErrNotAffected
	}
	return nil
}

func (p *poolDB) SetInterval(ctx context.Context, d time.Duration) (time.Duration, error) {
	var oldStr string
	const query = "UPDATE setting SET interval=$1 WHERE mine = TRUE RETURNING OLD.interval"
	err := p.QueryRow(ctx, query, d.String()).Scan(&oldStr)
	if err != nil {
		return 0, err
	}
	return time.ParseDuration(oldStr)
}

func (p *poolDB) ResizeWorker(ctx context.Context, workers uint) (uint, error) {
	var oldStr uint
	const query = "UPDATE setting SET workers=$1 WHERE mine = TRUE RETURNING old.workers"
	return oldStr, p.QueryRow(ctx, query, workers).Scan(&oldStr)
}
