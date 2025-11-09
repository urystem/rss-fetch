package storage

import (
	"context"
	"fmt"
	"rss/internal/configs"
	"rss/internal/ports/outbound"

	"github.com/jackc/pgx/v5/pgxpool"
)

type poolDB struct {
	*pgxpool.Pool
}

func InitDB(ctx context.Context, cfg configs.DBConfig) (outbound.PostgresInter, error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.GetUser(),
		cfg.GetPassword(),
		cfg.GetHostName(),
		cfg.GetPort(),
		cfg.GetDBName(),
		cfg.GetSSLMode(),
	)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &poolDB{pool}, pool.Ping(ctx)
}

func (p *poolDB) CloseDB() {
	p.Close()
}
