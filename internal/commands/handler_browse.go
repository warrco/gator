package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/warrco/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2

	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: browse <limit (optional)>")
	}

	if len(cmd.Args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("limit must be a number: %w", err)
		}
		limit = parsedLimit
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to retrieve posts for %v: %w", user, err)
	}

	for _, post := range posts {
		fmt.Println("Found post:")
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("  %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("======================================")
	}
	return nil
}
