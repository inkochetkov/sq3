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

	/*
		sql1, _ := db.Prepare("INSERT INTO `user` (`name`, `password`, `hash`) VALUES('test1', 'test2', 'test3')")
		tx, _ := db.Begin()
		_, _ = tx.Stmt(sql1).Exec()
		tx.Commit()
		defer sql1.Close()
	*/
	log.Println("Complite")
}
