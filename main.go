package main

import (
	"gregcsokas.hu/main/config"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if cfg.ServerEnv == "production" {
		// TODO add production related things here.
	}
}
