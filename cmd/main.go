package main

import (
	"fmt"
	"log"

	"github.com/jichong-tay/foodpanda-playlist-api/internal/config"
	"github.com/jichong-tay/foodpanda-playlist-api/internal/db"
)

const port = 3000

func main() {
	cfg := config.Init()

	dbClient, err := db.Init(cfg.DSN)
	if err != nil {
		log.Fatalf("db.Init failed: %v\n", err)
	}
	defer dbClient.Close()

	r := setupRoutes(dbClient)

	r.Run(fmt.Sprintf(":%d", port))
}
