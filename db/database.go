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
	CREATE TABLE IF NOT EXISTS barang (
		id TEXT PRIMARY KEY,
		nama TEXT NOT NULL,
		stok INTEGER NOT NULL
	);`

	_, err = db.Exec(queryTabel)
	if err != nil {
		log.Fatal("Error write to db:", err)
	}

	return db
}
