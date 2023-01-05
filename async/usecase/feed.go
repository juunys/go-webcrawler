package usecase

import (
	"time"

	"github.com/juunys/go-webcrawler/async/entity"
	"github.com/mmcdole/gofeed"
)

func ScrapeFeedPage(link, provider string) chan entity.Feed {
	fp := gofeed.NewParser()
	ch := make(chan entity.Feed)

	go func(l, p string) {
		defer close(ch)
		feed, _ := fp.ParseURL(l)

		for _, item := range feed.Items {
			ch <- entity.Feed{
				Title:       item.Title,
				Description: item.Description,
				Link:        item.Link,
				Provider:    p,
				Date:        item.Published,
				CreatedAt:   time.Now(),
			}
		}
	}(link, provider)

	return ch
}
