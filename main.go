package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/stg35/simple_bank/api"
	db "github.com/stg35/simple_bank/db/sqlc"
	"github.com/stg35/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot log config: ", err)
	}
	
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}