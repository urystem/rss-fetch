package cli

import (
	"flag"
	"fmt"
	"net/http"
	"rss/internal/domain"
)

func (c *cli) add() error {
	name := flag.String("name", "", "feed name")
	urlFlag := flag.String("url", "", "URL RSS")
	flag.Parse()
	// *name = strings.TrimSpace(*name)
	if len(*name) == 0 || len(*urlFlag) == 0 {
		return domain.ErrFlag
	}
	//check url
	resp, err := http.Head(*urlFlag)
	if err != nil {
		return err
	}
	if resp.StatusCode > 400 {
		return fmt.Errorf("URL is invalid, status: %s", resp.Status)
	}
	defer resp.Body.Close()
	return c.use.RssAdd(c.ctx, *name, resp.Request.URL.String())
}

func (c *cli) list() error {
	num := flag.Uint("num", 0, "URL RSS")
	flag.Parse()
	var wasNumSet bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "num" {
			wasNumSet = true
		}
	})
	if !wasNumSet {
		feeds, err := c.use.ListRssFeeds(c.ctx)
		if err != nil {
			return err
		}
		c.listFeedPrinter(feeds)
		return nil
	}
	if *num == 0 {
		return domain.ErrFlag
	}
	feeds, err := c.use.ListRssFeedsWithNum(c.ctx, *num)
	if err != nil {
		return err
	}
	c.listFeedPrinter(feeds)
	return nil
}

func (c *cli) listFeedPrinter(feeds []domain.Feed) {
	fmt.Println("Feeds:")
	for i, v := range feeds {
		fmt.Printf("\n%d. Name: %s\n", i+1, v.Name)
		fmt.Printf("   URL: %s\n", v.Url)
		fmt.Printf("   Added: %s\n", v.Created.String())
	}
}

func (c *cli) delete() error {
	name := flag.String("name", "", "feed name")
	flag.Parse()
	if len(*name) == 0 {
		return domain.ErrFlag
	}

	return c.use.DeleteRssFeed(c.ctx, *name)
}
