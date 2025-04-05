package commands

import (
	"context"
	"fmt"

	"github.com/warrco/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("failed to retrieve the user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
