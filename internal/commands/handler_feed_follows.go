package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/warrco/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}
	ctx := context.Background()
	url := cmd.Args[0]
	current_time := time.Now()
	id := uuid.New()

	feed, err := s.Db.GetFeedURL(ctx, url)
	if err != nil {
		return fmt.Errorf("failed to retrieve feed URL: %w", err)
	}

	feed_follows, err := s.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: current_time,
		UpdatedAt: current_time,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create new feed follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(feed_follows.UserName, feed_follows.FeedName)
	return nil
}

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("following command does not take any arguments")
	}
	ctx := context.Background()

	follows, err := s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to retrieve user feed follows: %w", err)
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, follow := range follows {
		fmt.Printf("* Following: 	%s\n", follow.FeedName)
	}
	return nil
}

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}
	url := cmd.Args[0]

	err := s.Db.DeleteFeedFollowsForUser(context.Background(), database.DeleteFeedFollowsForUserParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow the URL: %w", err)
	}
	fmt.Printf("Successfully unfollowed: %s\n", url)
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* Feed:			%s\n", feedname)
	fmt.Printf("* User: 		%s\n", username)
}
