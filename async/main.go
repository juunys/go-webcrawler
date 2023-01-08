package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/juunys/go-webcrawler/async/entity"
	"github.com/juunys/go-webcrawler/async/repository"
	"github.com/juunys/go-webcrawler/async/usecase"
	"github.com/juunys/go-webcrawler/common"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

func main() {
	feedRepository := repository.NewFeedRepository(repository.InitSqlite())
	var feedsSource []entity.FeedYml

	source, err := ioutil.ReadFile("db/feed.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(source, &feedsSource)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for {
		fmt.Printf("\n\nScraping url ...\n")

		var wg sync.WaitGroup
		chOut := make(chan entity.Feed)

		wg.Add(len(feedsSource))
		for _, feed := range feedsSource {
			go func(f entity.FeedYml, c chan entity.Feed) {
				defer wg.Done()
				usecase.ScrapeFeedPage(f.Link, f.Name, c)
			}(feed, chOut)
		}

		go func(ch chan entity.Feed) {
			for {
				feed, open := <-chOut
				if !open {
					break
				}
				feedRepository.InsertFeed(feed)
			}
		}(chOut)

		wg.Wait()
		close(chOut)

		fmt.Print("Sleeping for 1 hour ...\n")
		common.SleepBar()
	}
}
