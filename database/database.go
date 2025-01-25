package database

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

type Phone struct {
	Model string
	Brand string
	Price int
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./database/phones.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return db, nil
}

func SetupDB(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS phones (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        model TEXT NOT NULL,
        brand TEXT NOT NULL,
        price INTEGER NOT NULL
    );`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	return nil
}

func AddPhone(db *sql.DB, phone Phone) error {
	query := `INSERT INTO phones (model, brand, price) VALUES (?, ?, ?)`
	_, err := db.Exec(query, phone.Model, phone.Brand, phone.Price)
	if err != nil {
		return fmt.Errorf("failed to insert phone: %v", err)
	}
	return nil
}
