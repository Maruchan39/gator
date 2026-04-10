package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("name and url of the feed required")
	}

	if len(cmd.arguments) == 1 {
		return fmt.Errorf("url of the feed required")
	}

	name := cmd.arguments[0]
	url := cmd.arguments[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
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

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
