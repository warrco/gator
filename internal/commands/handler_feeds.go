package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/warrco/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	// Assign variables
	ctx := context.Background()
	name := cmd.Args[0]
	url := cmd.Args[1]
	feed_id := uuid.New()
	current_time := time.Now()
	follows_id := uuid.New()

	current_user, err := s.Db.GetUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to retrieve the user: %w", err)
	}

	feed, err := s.Db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        feed_id,
		CreatedAt: current_time,
		UpdatedAt: current_time,
		Name:      name,
		Url:       url,
		UserID:    current_user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to add feed to database: %w", err)
	}

	feed_follow, err := s.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        follows_id,
		CreatedAt: current_time,
		UpdatedAt: current_time,
		UserID:    current_user.ID,
		FeedID:    feed_id,
	})
	if err != nil {
		return fmt.Errorf("failed to create new feed follow: %w", err)
	}

	fmt.Println("Feed successfully created:")
	printFeed(feed, current_user)
	fmt.Println()
	fmt.Println("Feed followed successfully:")
	printFeedFollow(feed_follow.UserName, feed_follow.FeedName)
	fmt.Println("=========================================")
	return nil
}

func HandlerFeeds(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("feeds command does not take any arguments")
	}
	ctx := context.Background()

	feeds, err := s.Db.GetFeedsInfo(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* Name: 			%s\n", feed.Name_2)
		fmt.Printf("* URL:				%s\n", feed.Url)
		fmt.Printf("* User:				%s\n", feed.Name)
	}
	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID: 			%s\n", feed.ID)
	fmt.Printf("* Created: 		%s\n", feed.CreatedAt)
	fmt.Printf("* Updated: 		%s\n", feed.UpdatedAt)
	fmt.Printf("* Name: 		%s\n", feed.Name)
	fmt.Printf("* URL: 			%s\n", feed.Url)
	fmt.Printf("* User: 		%s\n", user.Name)
}
