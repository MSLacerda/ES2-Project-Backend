package database

import (
	"github.com/go-pg/pg"
	"log"
)

var db *pg.DB

func StartConn() {
	db = pg.Connect(&pg.Options{
		User: "postgres",
	})
}

func CloseConn() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func GetConn() *pg.DB {
	return db
}