package repository

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	e "github.com/juunys/go-webcrawler/async/entity"
)

type FeedRepository struct {
	db *sql.DB
}

func InitSqlite() *sql.DB {
	path := "db/development.sqlite3"
	file, err := os.Stat(path)
	if err != nil || file.Size() <= 0 {
		log.Println("Creating development.sqlite3...")
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()
		log.Println("development.sqlite3 created")
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err.Error())
	}

	const createFeedTable string = `
		CREATE TABLE IF NOT EXISTS feeds (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			provider TEXT,
			link TEXT UNIQUE,
			date TEXT,
			created_at DATETIME NOT NULL
		);`

	log.Println("Create feed table...")
	statement, err := db.Prepare(createFeedTable)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("feed table created")
	return db
}

func NewFeedRepository(db *sql.DB) *FeedRepository {
	return &FeedRepository{
		db: db,
	}
}

func (f *FeedRepository) InsertFeeds(feeds chan e.Feed) {
	insertSQL := `INSERT OR IGNORE INTO Feeds (title, description, provider, link, date, created_at ) VALUES(?,?,?,?,?,?)`
	stmt, err := f.db.Prepare(insertSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var wg sync.WaitGroup
	for feed := range feeds {
		wg.Add(1)
		go func(f e.Feed) {
			defer wg.Done()
			_, err = stmt.Exec(f.Title, f.Description, f.Provider, f.Link, f.Date, time.Now())
			if err != nil {
				log.Fatalln(err.Error())
			}
		}(feed)
	}
	wg.Wait()
}

func (f *FeedRepository) InsertFeed(feed e.Feed) bool {
	insertSQL := `INSERT OR IGNORE INTO Feeds (title, description, provider, link, date, created_at ) VALUES(?,?,?,?,?,?)`
	stmt, err := f.db.Prepare(insertSQL)
	if err != nil {
		log.Fatalln(err.Error())
		return false
	}
	_, err = stmt.Exec(feed.Title, feed.Description, feed.Provider, feed.Link, feed.Date, time.Now())
	if err != nil {
		log.Fatalln(err.Error())
		return false
	}

	return true
}
