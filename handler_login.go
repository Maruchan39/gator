package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("username required")
	}

	name := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user %s does not exist", name)
	}

	err = s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("username has been set")
	return nil
}
