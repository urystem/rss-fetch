package cli

import (
	"flag"
	"fmt"
	"rss/internal/domain"
)

func (c *cli) setInterval() error {
	dur := flag.Duration("duration", 0, "")
	flag.Parse()
	if *dur == 0 {
		return domain.ErrFlag
	}
	old, err := c.use.SetInterval(c.ctx, *dur)
	if err != nil {
		return err
	}
	fmt.Printf("Interval of fetching feeds changed from %s minutes to %s minutes\n", old.String(), dur.String())
	return nil
}

func (c *cli) setWorker() error {
	count := flag.Uint("count", 0, "")
	flag.Parse()
	if *count == 0 {
		return domain.ErrFlag
	}
	old, err := c.use.ResizeWorker(c.ctx, *count)
	if err != nil {
		return err
	}
	fmt.Printf("Number of workers changed from %d to %d\n", old, *count)
	return nil
}
