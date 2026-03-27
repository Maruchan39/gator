package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("username required")
	}

	name := cmd.arguments[0]
	user_id := uuid.New()
	created_at := time.Now()
	updated_at := time.Now()

	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("user %s already exists", name)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        user_id,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User was created successfully!")
	fmt.Printf("User data: ID=%s, Name=%s, CreatedAt=%s, UpdatedAt=%s\n",
		user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

	return nil
}
