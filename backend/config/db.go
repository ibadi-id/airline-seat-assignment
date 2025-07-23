package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./vouchers.db")
	if err != nil {
		log.Fatal(err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS vouchers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		crew_name TEXT,
		crew_id TEXT,
		flight_number TEXT,
		flight_date TEXT,
		aircraft_type TEXT,
		seat1 TEXT,
		seat2 TEXT,
		seat3 TEXT,
		created_at TEXT
	);
	`
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
