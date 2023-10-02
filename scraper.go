package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

func startScraping(db *database.Queries, conncurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", conncurrency, timeBetweenRequest)
	ticker := time.NewTimer(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(conncurrency))
		if err != nil {
			log.Println("Error while fetching the feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error while marking the feed as fetched: ", err)
		return
	}

	rssFeed, err := urlToRssFeed(feed.Url)
	if err != nil {
		log.Println("Error while fetching the feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Post found:", item.Title, "on feed", feed.Name)
	}
	log.Printf("Feeds %s fetched, Total number of post found %d", feed.Name, len(rssFeed.Channel.Item))
}
