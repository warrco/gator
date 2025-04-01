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
	feedID := uuid.New()
	current_time := time.Now()

	current_user, err := s.Db.GetUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to retrieve the user: %w", err)
	}

	feed, err := s.Db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        feedID,
		CreatedAt: current_time,
		UpdatedAt: current_time,
		Name:      name,
		Url:       url,
		UserID:    current_user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to add feed to database: %w", err)
	}

	fmt.Println("Feed successfully created:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=========================================")
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID: 			%s\n", feed.ID)
	fmt.Printf("* Created: 		%s\n", feed.CreatedAt)
	fmt.Printf("* Updated: 		%s\n", feed.UpdatedAt)
	fmt.Printf("* Name: 		%s\n", feed.Name)
	fmt.Printf("* URL: 			%s\n", feed.Url)
	fmt.Printf("* UserID: 		%s\n", feed.UserID)
}
