package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/warrco/gator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	userID := uuid.New()
	currentTime := time.Now()
	ctx := context.Background()

	user, err := s.Db.CreateUser(ctx, database.CreateUserParams{
		ID:        userID,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      name,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			fmt.Printf("Internal log: unique_violation while trying to register: %v\n", err)
			return fmt.Errorf("registration error")
		}

		return fmt.Errorf("failed to create the user: %v", err)
	}

	cfg := s.Config

	err = cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

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
	name := cmd.Args[0]

	ctx := context.Background()

	_, err := s.Db.GetUser(ctx, name)
	if err != nil {
		fmt.Printf("could not find user: %v\n", err)
		return fmt.Errorf("unable to login")
	}

	err = cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}

	fmt.Printf("Username has been set to: %s\n", name)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID: 		%v\n", user.ID)
	fmt.Printf(" * Name		%v\n", user.Name)
}
