package main

import (
	"database/sql"
	"go_mindmap/core/usecases"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func setUpDB() *sql.DB { //Roman fragen nach *xyz logik
	db, err := sql.Open("sqlite3", "./mindmap.db")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	creationStatement := `
	CREATE TABLE IF NOT EXISTS mmaps (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		node TEXT 
	);
	`

	_, err = db.Exec(creationStatement)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'mmap' angelegt")
	return db
}

func commitMmapToDB(f *sql.DB, g usecases.Node) {
	query := "INSERT INTO mmaps(node) VALUES(?)"
	_, err := f.Exec(query, g.Content)
	if err != nil {
		log.Printf("failure to insert node: %v", err)
		return
	}

	log.Println("Node saved in DB")
}
