package usecase

import (
	"time"

	"github.com/juunys/go-webcrawler/async/entity"
	"github.com/mmcdole/gofeed"
)

func ScrapeFeedPage(link, provider string, out chan entity.Feed) chan entity.Feed {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(link)

	for _, item := range feed.Items {
		out <- entity.Feed{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			Provider:    provider,
			Date:        item.Published,
			CreatedAt:   time.Now(),
		}
	}

	return out
}
