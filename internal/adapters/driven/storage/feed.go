package storage

import (
	"context"
	"errors"
	"rss/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
)

func (p *poolDB) RssAdd(ctx context.Context, name, url string) error {
	const query = `
        INSERT INTO feeds (name, url)
        VALUES ($1, $2)
    `
	_, err := p.Exec(ctx, query, name, url)
	if err != nil {
		// проверяем уникальность имени
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok && pgErr.Code == "23505" {
			return domain.ErrConflict
		}
	}
	return err
}
func (p *poolDB) ListRssFeeds(ctx context.Context) ([]domain.Feed, error) {
	const query = `
        SELECT id, name, url, created_at, updated_at
        FROM feeds
        ORDER BY created_at`

	rows, err := p.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feeds []domain.Feed
	for rows.Next() {
		var f domain.Feed
		err = rows.Scan(
			&f.ID,
			&f.Name,
			&f.Url,
			&f.Created,
			&f.Updated,
		)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, f)
	}

	// Проверка на ошибки после окончания чтения всех строк
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return feeds, nil
}

func (p *poolDB) ListRssFeedsWithNum(ctx context.Context, n uint) ([]domain.Feed, error) {
	const query = `
        SELECT id, name, url, created_at, updated_at
        FROM feeds
        ORDER BY created_at
        LIMIT $1`

	rows, err := p.Query(ctx, query, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feeds := make([]domain.Feed, 0, n)
	for rows.Next() {
		var f domain.Feed
		err = rows.Scan(
			&f.ID,
			&f.Name,
			&f.Url,
			&f.Created,
			&f.Updated,
		)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, f)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return feeds, nil
}

func (p *poolDB) DeleteRssFeed(ctx context.Context, name string) error {
	res, err := p.Exec(ctx, `DELETE FROM feeds WHERE name = $1`, name)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return domain.ErrNotAffected
	}

	return nil
}
func (p *poolDB) ListRssWithLastUpdated(ctx context.Context, n uint) ([]domain.Feed, error) {
	const query = `
        SELECT id, name, url, created_at, updated_at
        FROM feeds
        ORDER BY updated_at ASC
        LIMIT $1`

	rows, err := p.Query(ctx, query, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feeds := make([]domain.Feed, 0, n)
	for rows.Next() {
		var f domain.Feed
		if err := rows.Scan(&f.ID, &f.Name, &f.Url, &f.Created, &f.Updated); err != nil {
			return nil, err
		}
		feeds = append(feeds, f)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return feeds, nil
}

// базар жоқ
func (p *poolDB) UpdateFeedsLastUpdates(ctx context.Context, feeds []string) error {
	if len(feeds) == 0 {
		return nil
	}
	const query = `
        UPDATE feeds
        SET updated_at = NOW()
        WHERE id = ANY($1)`

	_, err := p.Exec(ctx, query, feeds)
	return err
}
