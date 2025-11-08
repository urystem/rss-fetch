package storage

import (
	"context"
	"fmt"
	"time"
)

func (p *poolDB) Starter(ctx context.Context) error {
	res, err := p.Exec(ctx, `
        UPDATE setting
        SET is_running = TRUE
        WHERE mine = TRUE
    `)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("not found: no setting with mine = TRUE")
	}
	return nil
}

func (p *poolDB) Shutdown(ctx context.Context) error {
	res, err := p.Exec(ctx, `
        UPDATE setting
        SET is_running = FALSE
        WHERE mine = TRUE
    `)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("not found: no setting with mine = TRUE")
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
