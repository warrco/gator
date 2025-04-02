package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/warrco/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}
	ctx := context.Background()
	url := cmd.Args[0]
	current_time := time.Now()
	id := uuid.New()

	current_user, err := s.Db.GetUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to retrieve the user: %w", err)
	}

	feed, err := s.Db.GetFeedURL(ctx, url)
	if err != nil {
		return fmt.Errorf("failed to retrieve feed URL: %w", err)
	}

	feed_follows, err := s.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: current_time,
		UpdatedAt: current_time,
		UserID:    current_user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create new feed follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(feed_follows.UserName, feed_follows.FeedName)
	return nil
}

func HandlerFollowing(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("following command does not take any arguments")
	}
	ctx := context.Background()

	current_user, err := s.Db.GetUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to retrieve the user: %w", err)
	}

	follows, err := s.Db.GetFeedFollowsForUser(ctx, current_user.ID)
	if err != nil {
		return fmt.Errorf("failed to retrieve user feed follows: %w", err)
	}

	fmt.Printf("Feed follows for user %s:\n", current_user.Name)
	for _, follow := range follows {
		fmt.Printf("* Following: 	%s\n", follow.FeedName)
	}
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* Feed:			%s\n", feedname)
	fmt.Printf("* User: 		%s\n", username)
}
