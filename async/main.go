package main

import (
	"fmt"
	"sync"

	"github.com/juunys/go-webcrawler/async/repository"
	"github.com/juunys/go-webcrawler/async/usecase"
	"github.com/juunys/go-webcrawler/common"
	"github.com/juunys/go-webcrawler/entity"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app, err := common.NewApp()
	if err != nil {
		panic(err.Error())
	}

	db, err := repository.InitSqlite()
	if err != nil {
		panic(err.Error())
	}
	feedRepository := repository.NewFeedRepository(db)

	for {
		fmt.Printf("\n\nScraping url ...\n")

		var wg sync.WaitGroup
		var m sync.Mutex
		chOut := make(chan entity.Feed)
		scrappedCount := 0

		wg.Add(len(app.Source))
		for _, feed := range app.Source {
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
