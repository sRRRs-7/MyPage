package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sRRRs-7/MyPage/api"
	db "github.com/sRRRs-7/MyPage/db/sqlc"
	"github.com/sRRRs-7/MyPage/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(" Error loading config ", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal(" Error connect database ", err)
	}

	store := db.NewStore(conn)

	runGinServer(config, store)
}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.HTTP_SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}