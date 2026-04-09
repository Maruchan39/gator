package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	username := s.config.CurrentUserName

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollows {

		fmt.Printf("* %s\n", feedFollow.FeedName)

	}

	return nil
}
