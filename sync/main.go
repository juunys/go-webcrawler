package main

import (
	"fmt"

	"github.com/juunys/go-webcrawler/common"
	"github.com/juunys/go-webcrawler/entity"
	"github.com/juunys/go-webcrawler/sync/repository"
	"github.com/juunys/go-webcrawler/sync/usecase"
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
		var feeds = []*entity.Feed{}

		for _, feed := range app.Source {
			generatedFeed := usecase.ScrapeFeedPage(feed.Link, feed.Name)
			feeds = append(feeds, generatedFeed...)
		}

		feedRepository.InsertFeeds(feeds)

		fmt.Printf("Sleeping for 1 hour ...\n")
		common.SleepBar()
	}
}
