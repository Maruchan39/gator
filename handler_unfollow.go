package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("feed url required")
	}

	feedUrl := cmd.arguments[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("feed unfollowed")
	return nil
}
