package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "./goedang.db")
	if err != nil {
		log.Fatal("Error connect database:", err)
	}

	queryTabel := `
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		stock INTEGER NOT NULL,
		price INTEGER NOT NULL,
		time_stamp TEXT DEFAULT (DATETIME('now', 'localtime'))
	);`

	_, err = db.Exec(queryTabel)
	if err != nil {
		log.Fatal("Error write to db:", err)
	}

	return db
}
