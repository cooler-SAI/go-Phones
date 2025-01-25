package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"os"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

// WelcomePage отображает главную страницу
func WelcomePage(w http.ResponseWriter, _ *http.Request) {
	// Читаем HTML из файла
	htmlContent, err := os.ReadFile("html/welcome.html")
	if err != nil {
		http.Error(w, "Failed to load HTML template", http.StatusInternalServerError)
		return
	}

	// Отправляем HTML клиенту
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = w.Write(htmlContent)
	if err != nil {
		return
	}
}

// PhoneListPage отображает список телефонов
func PhoneListPage(w http.ResponseWriter, _ *http.Request) {
	// Запрос данных из базы
	rows, err := DB.Query("SELECT model, brand, price FROM phones")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	// Собираем данные в срез
	var phones []struct {
		Model string
		Brand string
		Price float64
	}
	for rows.Next() {
		var model, brand string
		var price float64
		if err := rows.Scan(&model, &brand, &price); err != nil {
			http.Error(w, "Failed to read data", http.StatusInternalServerError)
			return
		}
		phones = append(phones, struct {
			Model string
			Brand string
			Price float64
		}{model, brand, price})
	}

	// Читаем HTML-шаблон из файла
	tmpl, err := template.ParseFiles("html/phone_list.html")
	if err != nil {
		http.Error(w, "Failed to load HTML template", http.StatusInternalServerError)
		return
	}

	// Применяем шаблон и отправляем результат клиенту
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, phones); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
