package main

import (
	"fmt"
	"sync"

	"github.com/juunys/go-webcrawler/async/entity"
	"github.com/juunys/go-webcrawler/async/repository"
	"github.com/juunys/go-webcrawler/async/usecase"
	"github.com/juunys/go-webcrawler/common"
	_ "github.com/mattn/go-sqlite3"
)

func fanIn(inputs ...chan entity.Feed) chan entity.Feed {
	var wg sync.WaitGroup
	out := make(chan entity.Feed)

	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(ch <-chan entity.Feed) {
			for {
				feed, ok := <-ch

				if !ok {
					wg.Done()
					break
				}

				out <- feed
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	feedRepository := repository.NewFeedRepository(repository.InitSqlite())

	for {
		fmt.Print("\n\n")
		fmt.Println("Scraping url ...")
		fmt.Print("\n")

		f1 := usecase.ScrapeFeedPage("https://catracalivre.com.br/feed/", "catraca_livre")
		f2 := usecase.ScrapeFeedPage("https://www.infomoney.com.br/feed/", "infomoney")
		f3 := usecase.ScrapeFeedPage("https://forbes.com.br/feed/", "forbes")
		f4 := usecase.ScrapeFeedPage("https://www.cnnbrasil.com.br/feed/", "cnn")
		f5 := usecase.ScrapeFeedPage("https://www.moneytimes.com.br/feed/", "moneytimes")

		out := fanIn(f1, f2, f3, f4, f5)
		feedRepository.InsertFeeds(out)

		fmt.Println("Sleeping for 1 hour ...")
		fmt.Println()
		common.SleepBar()
	}
}
