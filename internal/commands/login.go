package commands

import (
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	// If arguments slice is empty, return an error
	if len(cmd.Args) == 0 {
		return fmt.Errorf("must enter a valid username")
	}
	//Handle too many usernames
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usernames must not contain spaces")
	}

	//Set the username
	cfg := s.Config

	err := cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Username has been set to: %s\n", cmd.Args[0])
	return nil
}
