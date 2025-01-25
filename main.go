package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"go-Phones/database"
	"go-Phones/handlers"
	"net/http"
	"os"
	"strconv"
	"strings"

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

	// Установим базу данных для обработчиков
	handlers.SetDB(db)

	// Запуск веб-сервера в отдельной горутине
	go func() {
		log.Info().Msg("Welcome to the Phones Web Server Application!")
		log.Info().Msg("Server is running!")
		log.Info().Msg("Main page: http://localhost:8080")
		log.Info().Msg("Phone list: http://localhost:8080/phones")

		fs := http.FileServer(http.Dir("static"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
		http.HandleFunc("/", handlers.WelcomePage)         // Главная страница
		http.HandleFunc("/phones", handlers.PhoneListPage) // Страница со списком телефонов

		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()

	// Обработка команд из консоли
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter commands (e.g., .add-phone <model> <brand> <price>):")
	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case ".add-phone":
			if len(args) != 4 {
				fmt.Println("Usage: .add-phone <model> <brand> <price>")
				continue
			}

			price, err := strconv.Atoi(args[3])
			if err != nil {
				fmt.Println("Invalid price:", err)
				continue
			}

			phone := database.Phone{
				Model: args[1],
				Brand: args[2],
				Price: price,
			}

			if err := database.AddPhone(db, phone); err != nil {
				fmt.Println("Failed to add phone:", err)
			} else {
				fmt.Println("Phone added successfully!")
			}

		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
