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
	feedRepository := repository.NewFeedRepository(repository.InitSqlite())

	providers := []string{"catraca_livre", "infomoney", "forbes", "cnn", "moneytimes"}
	urlFeed := []string{"https://catracalivre.com.br/feed/", "https://www.infomoney.com.br/feed/", "https://forbes.com.br/feed/", "https://www.cnnbrasil.com.br/feed/", "https://www.moneytimes.com.br/feed/"}

	for {
		fmt.Print("\n\n")
		fmt.Println("Scraping url ...")
		fmt.Print("\n")
		var feeds = []*entity.Feed{}

		for index, url := range urlFeed {
			generatedFeed := usecase.ScrapeFeedPage(url, providers[index])
			feeds = append(feeds, generatedFeed...)
		}

		feedRepository.InsertFeeds(feeds)

		fmt.Println("Sleeping for 1 hour ...")
		fmt.Println()
		common.SleepBar()
	}
}
