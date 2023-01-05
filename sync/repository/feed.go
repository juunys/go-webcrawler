package repository

import (
	e "github.com/juunys/go-webcrawler/entity"
)

type HabitRepository interface {
	InsertFeeds(feed []*e.Feed) bool
	InsertFeed(feed e.Feed) bool
}
