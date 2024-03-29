// package main is the entry point for the playlist service.
package main

import (
	"database/sql"
	"log"

	"github.com/jichong-tay/playlist-api/api"
	db "github.com/jichong-tay/playlist-api/db/sqlc"
	"github.com/jichong-tay/playlist-api/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
