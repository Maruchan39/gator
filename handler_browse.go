package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {
	limit := 2

	if len(cmd.arguments) > 0 {
		i, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return err
		}

		limit = i
	}

	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("* %s\n", post.Title.String)
		fmt.Printf("* %s\n", post.Description.String)
	}

	return nil
}
