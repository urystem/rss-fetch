package domain

import "time"

type Setting struct {
	IsRunning bool
	Worker    int
	Interval  time.Duration
}
