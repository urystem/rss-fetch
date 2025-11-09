package domain

import (
	"time"
)

type Feed struct {
	FeedForGetReq
	Name    string
	Created time.Time
	Updated time.Time
}

type FeedForGetReq struct {
	ID  string //for insert
	Url string //for get
}
