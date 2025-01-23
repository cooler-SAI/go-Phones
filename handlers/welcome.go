package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

func WelcomePage(w http.ResponseWriter, _ *http.Request) {
	rows, err := DB.Query("SELECT name, brand, price FROM phones")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var name, brand string
		var price float64
		err := rows.Scan(&name, &brand, &price)
		if err != nil {
			return
		}
		write, err := w.Write([]byte(name + " - " + brand + " - " +
			fmt.Sprintf("%.2f", price) + "\n"))
		if err != nil {
			return
		}
		fmt.Printf("Written %d bytes\n", write)

	}
}
