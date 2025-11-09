package cli

import (
	"flag"
	"fmt"
	"rss/internal/domain"
)

func (c *cli) showArticles() error {
	name := flag.String("feed-name", "", "feed name")
	num := flag.Uint("num", 3, "number of articles")
	flag.Parse()
	arc, err := c.use.ShowArticles(c.ctx, *name, *num)
	if err != nil {
		return err
	}
	c.articlePrinter(arc)
	return nil
}

func (c *cli) articlePrinter(articles []domain.Article) {
	if len(articles) == 0 {
		fmt.Println("No articles found.")
		return
	}

	fmt.Println()
	for i, a := range articles {
		fmt.Printf("%d. [%s] %s\n",
			i+1,
			a.PublishedAt.Format("2006-01-02"),
			a.Title,
		)
		fmt.Printf("   %s\n", a.Link)
		fmt.Println()
	}
}
