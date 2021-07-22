package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "init/bd.sq3")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	log.Println("Test complite")
}
