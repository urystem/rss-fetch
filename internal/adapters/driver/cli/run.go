package cli

import "rss/internal/ports/inbound"

type cli struct {
	tick inbound.TickController
	work inbound.WorkerForCLI
}

func BuildCli() any {
	return nil
}
