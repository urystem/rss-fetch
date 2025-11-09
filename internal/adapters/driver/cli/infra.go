package cli

import (
	"context"
	"fmt"
	"os"
	"rss/internal/domain"
	"rss/internal/ports/inbound"
)

type cli struct {
	ctx context.Context
	use inbound.UseCasePsql
}

func BuildCli(use inbound.UseCasePsql) inbound.CliInter {
	return &cli{
		use: use,
	}
}

func (c *cli) Run(ctx context.Context) error {
	if c.ctx != nil {
		return fmt.Errorf("already used cli")
	}
	c.ctx = ctx
	if len(os.Args) < 2 {
		return domain.ErrFlag
	}
	//for flag.Parse we need to cut the first element of os.Args)
	command := os.Args[1]
	os.Args = os.Args[1:]
	switch command {
	case "fetch":
		return c.use.Starter(ctx)
	case "add":
		return c.add()
	case "set-interval":
		return c.setInterval()
	case "set-workers":
		return c.setWorker()
	case "list":
		return c.list()
	case "delete":
		return c.delete()
	case "articles":
		return c.showArticles()
	case "stop-fetch":
		return c.use.Stopper(ctx)
	default:
		return domain.ErrFlag
	}
}
