package main

import (
	"context"
	"fmt"
	"time"

	"gator/internal/database"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.arguments) == 0 {
		return fmt.Errorf("feed url required")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	id := uuid.New()
	user_id := user.ID
	feed_id := feed.ID
	created_at := time.Now()
	updated_at := time.Now()

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		UserID:    user_id,
		FeedID:    feed_id,
	})
	if err != nil {
		return err
	}

	fmt.Printf("* %s\n", feedFollow.FeedName)
	fmt.Printf("* %s\n", user.Name)

	return nil
}
