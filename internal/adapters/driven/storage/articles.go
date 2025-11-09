package storage

import (
	"context"
	"rss/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (p *poolDB) ShowArticles(ctx context.Context, name string, n uint) ([]domain.Article, error) {
	const query = `
        UPDATE articles a
        SET last_seen = NOW()
        FROM feeds f
        WHERE a.feed_id = f.id
          AND f.name = $1
        ORDER BY a.last_seen ASC NULLS FIRST, a.published_at DESC
        LIMIT $2
        RETURNING a.id, a.title, a.link, a.description, a.published_at, a.created_at, a.last_seen`

	rows, err := p.Query(ctx, query, name, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]domain.Article, 0, n)
	for rows.Next() {
		var a domain.Article
		if err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Link,
			&a.Description,
			&a.PublishedAt,
			&a.CreatedAt,
			&a.LastSeen,
		); err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return articles, nil
}

func (p *poolDB) InsertArticles(ctx context.Context, feedID string, articles []domain.RSSItem) error {
	if len(articles) == 0 {
		return nil
	}

	batch := &pgx.Batch{}
	for _, a := range articles {
		batch.Queue(`
            INSERT INTO articles (title, link, description, published_at, feed_id)
            VALUES ($1, $2, $3, $4, $5)
            ON CONFLICT (link) DO NOTHING`,
			a.Title, a.Link, a.Description, a.PubDate, feedID)
	}

	br := p.SendBatch(ctx, batch)
	defer br.Close()

	// Проверяем результат каждой вставки
	for range articles {
		_, err := br.Exec()
		if err != nil {
			return err
		}
	}
	return nil
}
