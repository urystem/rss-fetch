package cli

import (
	"context"
	"os"
	"rss/internal/domain"
	"rss/internal/ports/inbound"
)

type cli struct {
	use inbound.UseCasePsql
}

func BuildCli(use inbound.UseCasePsql) inbound.CliInter {
	return &cli{
		use: use,
	}
}

func (c *cli) Run(ctx context.Context) error {
	if len(os.Args) < 2 {
		return domain.ErrHelp
	}
	switch os.Args[1] {
	case "fetch":
		return c.use.Starter(ctx)
	case "add":
		return c.add()
	case "set-interval":
		return nil
	case "set-workers":
		return nil
	case "list":
		return nil
	case "delete":
		return nil
	case "articles":
		return nil
	case "stop-fetch":
		return c.use.Stopper(ctx)
	default:
		return domain.ErrHelp
	}
}

func (c *cli) add() error {
	return nil
}
