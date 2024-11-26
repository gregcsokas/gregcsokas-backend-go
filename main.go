package main

import (
	"github.com/gin-gonic/gin"
	"gregcsokas.hu/main/config"
	"gregcsokas.hu/main/db"
	"gregcsokas.hu/main/modules/auth"
	"gregcsokas.hu/main/modules/newsletter"
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
	database := db.InitDB(
		cfg.GetDSN(),
	)

	err = database.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	authModule := auth.NewModule(database)
	newsletterModule := newsletter.NewModule(database)

	api := server.Group("/api")
	v1 := api.Group("/v1")

	auth.RegisterRoutes(v1, authModule)
	newsletter.RegisterRoutes(v1, newsletterModule)

	err = server.Run(":8060")
	if err != nil {
		log.Fatal(err)
	}
}
