package main

import (
	"database/sql"
	"go-Phones/commands"
	"go-Phones/database"
	"go-Phones/handlers"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

func main() {
	// Настройка Zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Инициализация базы данных
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close database")
		}
	}(db)

	// Настройка базы данных (создание таблицы, если её нет)
	if err := database.SetupDB(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to set up database")
	}

	// Проверяем аргументы командной строки
	if len(os.Args) > 1 {
		// Если переданы аргументы, запускаем CLI
		if err := commands.Execute(db); err != nil {
			log.Fatal().Err(err).Msg("CLI execution error")
		}
		return
	}

	// Если аргументов нет, запускаем веб-сервер
	log.Info().Msg("Welcome to the Phones Web Server Application!")
	log.Info().Msg("Server is running!")
	log.Info().Msg("Main page: http://localhost:8080")
	log.Info().Msg("Phone list: http://localhost:8080/phones")

	// Установим базу данных для обработчиков
	handlers.SetDB(db)

	// Запуск веб-сервера
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.WelcomePage)         // Главная страница
	http.HandleFunc("/phones", handlers.PhoneListPage) // Страница со списком телефонов

	// Запуск сервера
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().Err(err).Msg("Server failed")
	}
}
