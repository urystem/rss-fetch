package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"rss/internal/domain"
	"rss/internal/ports/outbound"
	"time"
)

type client struct {
	http.Client
}

func BuildRss(timeout time.Duration) outbound.RssHttp {
	return &client{
		http.Client{Timeout: timeout},
	}
}

func (c *client) GetRss(ctx context.Context, url string) ([]domain.RSSItem, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	var feed domain.RSSFeed
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return feed.Channel.Item, nil
}
