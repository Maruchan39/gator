package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("martynas")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
