package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	username := s.config.CurrentUserName

	if len(cmd.arguments) == 0 {
		return fmt.Errorf("name and url of the feed required")
	}

	if len(cmd.arguments) == 1 {
		return fmt.Errorf("url of the feed required")
	}

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	name := cmd.arguments[0]
	url := cmd.arguments[1]
	id := uuid.New()
	user_id := user.ID
	created_at := time.Now()
	updated_at := time.Now()

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        id,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		Name:      name,
		Url:       url,
		UserID:    user_id,
	})

	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	fmt.Printf(" * ID:        %s\n", feed.ID)
	fmt.Printf(" * Name:      %s\n", feed.Name)
	fmt.Printf(" * URL:       %s\n", feed.Url)
	fmt.Printf(" * UserID:    %s\n", feed.UserID)
	fmt.Printf(" * CreatedAt: %s\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt: %s\n", feed.UpdatedAt)

	return nil
}
