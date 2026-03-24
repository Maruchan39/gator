package main

import (
	"log"
	"os"

	"gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := state{config: &cfg}
	cmds := commands{handlers: map[string]func(*state, command) error{}}

	cmds.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("cli argument required")
	}

	cmd := command{
		name:      args[1],
		arguments: args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
