package commands

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("reset command does not take any arguments")
	}
	ctx := context.Background()

	err := s.Db.DeleteUser(ctx)
	if err != nil {
		return fmt.Errorf("unable to reset user records: %w", err)
	}
	fmt.Println("Successfully reset users")
	return nil
}
