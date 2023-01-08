package entity

import "time"

type Feed struct {
	Title       string
	Description string
	Provider    string
	Link        string
	Date        string
	CreatedAt   time.Time
}

type FeedYml struct {
	Name string
	Link string
}
