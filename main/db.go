package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func setUpDB() {
	db, err := sql.Open("sqlite3", "./mindmap.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS mmaps (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		node TEXT 
	);
	`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'mmap' angelegt")

}
