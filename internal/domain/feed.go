package domain

import (
	"time"
)

type Feed struct {
	ID      string
	Name    string
	Url     string
	Created time.Time
	Updated time.Time
}
