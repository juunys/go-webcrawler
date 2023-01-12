package usecase

import (
	"time"

	e "github.com/juunys/go-webcrawler/entity"
	"github.com/mmcdole/gofeed"
)

func ScrapeFeedPage(link, provider string) []*e.Feed {
	fp := gofeed.NewParser()
	feeds := []*e.Feed{}
	feed, _ := fp.ParseURL(link)
	for _, item := range feed.Items {
		feedParse := e.Feed{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			Provider:    provider,
			Date:        item.Published,
			CreatedAt:   time.Now(),
		}
		feeds = append(feeds, &feedParse)
	}

	return feeds
}
