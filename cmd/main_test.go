package main

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func TestConnect(t *testing.T) {
	db, err := sql.Open("sqlite3", "init/bd.sq3")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	log.Println("Test complite")

}
