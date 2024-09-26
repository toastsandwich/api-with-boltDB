package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/toastsandwich/restraunt-api-system/cmd/api"
)

func main() {
	// create db
	db, err := bolt.Open("db/restraunt.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	// create new instance for api
	api, err := api.NewAPIServer(":8080", db)
	if err != nil {
		log.Fatal(err)
	}
	// start api server
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
