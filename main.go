package main

import (
	"github.com/gin-gonic/gin"
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
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()

	err = server.Run(":8060")
	if err != nil {
		log.Fatal(err)
	}
}
