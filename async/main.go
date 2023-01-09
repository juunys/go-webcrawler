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
	var feedsSource []entity.FeedYml
	source, err := ioutil.ReadFile("db/feed.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(source, &feedsSource)
	if err != nil {
		log.Panicf("error: %v", err)
	}

	feedRepository := repository.NewFeedRepository(repository.InitSqlite())

	for {
		fmt.Printf("\n\nScraping url ...\n")

		var wg sync.WaitGroup
		var m sync.Mutex
		chOut := make(chan entity.Feed)
		scrappedCount := 0

		wg.Add(len(feedsSource))
		for _, feed := range feedsSource {
			go usecase.ScrapeFeedPage(feed.Link, feed.Name, chOut, &wg)
		}

		go func() {
			for {
				feed, open := <-chOut
				if !open {
					break
				}
				feedRepository.InsertFeed(feed)
				m.Lock()
				scrappedCount++
				m.Unlock()
			}
		}()

		wg.Wait()
		close(chOut)
		m.Lock()
		fmt.Printf("Scrapped %d items\n\n", scrappedCount)
		m.Unlock()

		fmt.Print("Sleeping for 1 hour ...\n")
		common.SleepBar()
	}
}
