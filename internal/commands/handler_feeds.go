package commands

import (
	"context"
	"fmt"
)

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
		fmt.Printf("* Created By:			%s\n", feed.Name)
	}
	return nil
}
