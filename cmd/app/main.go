package main

import (
	"context"
	"frontend_go/catalog"
	"frontend_go/events"
	"frontend_go/server"
	"frontend_go/store"
	"log"
	"os"
)

func main() {
	path := os.Getenv("BOARD_DB")
	if path == "" {
		path = "boards.db"
	}
	db, e := store.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()
	c := catalog.New(db)
	d := events.NewDispatcher(db, nil)
	log.Fatal(server.Run(context.Background(), ":8080", server.New(c, d).Handler()))
}
