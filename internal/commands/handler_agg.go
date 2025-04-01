package commands

import (
	"context"
	"fmt"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("agg command does not take any arguments")
	}

	ctx := context.Background()
	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := FetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("Title: %s\n", feed.Channel.Title)
	fmt.Printf("Description: %s\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("Item: %s (%s)\n", item.Title, item.Link)
	}
	return nil
}
