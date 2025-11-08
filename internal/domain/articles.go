package domain

import (
	"time"
)

type Article struct {
	ID          string
	Title       string
	Link        string
	Description string
	PublishedAt time.Time
	CreatedAt   time.Time
	LastSeen    *time.Time
}
