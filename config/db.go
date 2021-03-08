package config

import (
	"log"
	"os"

	controllers "github.com/bhargavameesala0420/go_db/controllers"
	"github.com/go-pg/pg"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "root",
		Addr:     "localhost:5432",
		Database: "test_connect",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
