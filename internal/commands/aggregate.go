package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/warrco/gator/internal/database"
)

func ScrapeFeeds(s *State) {
	next_feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("could not retrieve next feed to fetch")
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.Db, next_feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Could not mark the feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Could not collect feed %s: %v", feed.Name, err)
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
