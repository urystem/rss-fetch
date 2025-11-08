package worker

import "time"

type worker struct {
	countWorker uint
	interval    time.Duration
}

func BuildWorker()