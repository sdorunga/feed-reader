package main

import (
	"github.com/boltdb/bolt"

	"feed-reader/app"
	"feed-reader/feedlist"
	"feed-reader/fetcher"

	"log"
	"net/http"
)

func main() {
	db, err := bolt.Open("feed-reader.db", 0600, nil)
	if err != nil {
		log.Fatalf("Error opening db: %v", err)
	}
	defer db.Close()

	listStore, err := feedlist.NewFeedListStore(db)
	if err != nil {
		log.Fatalf("Error initialising db: %v", err)
	}
	cachedFetcher := fetcher.NewCachedFetcher(fetcher.NewFeedFetcher())

	log.Println("Running on localhost:8080")
	http.ListenAndServe(":8080", app.InitRouter(listStore, cachedFetcher))
}
