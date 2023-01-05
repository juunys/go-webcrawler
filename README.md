# Go RSS Webcrawler
This project is a rss webcrawler in go, that search for feed post every hour and store the data in the database.

## Requirements

| Name | Version | Notes | Mandatory
|------|---------|---------|---------|
| [golang](https://golang.org/dl/) | >= go1.18 | Main programming language | true
| [make](https://www.gnu.org/software/make/) | Run shortcuts | n/a | false

### Dependencies
The mandatory tools/libs to run this webcrawler

| Name | Notes | command
|------|---------|---------|
| gofeed | Used to rss webcrawler | `go get github.com/mmcdole/gofeed`


## Providers

| Name | Version | Notes | command
|------|---------|---------|---------|
| [sqlite]| any stable version | in this repo, I'm using sqlite3 | `go get github.com/mattn/go-sqlite3`

# Usage
Into root folder run:

```
make s                  # Run script on sync mode
```

or

```
make a                  # Run script with goroutine
```