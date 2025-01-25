package commands

import (
	"database/sql"
	"errors"
	"fmt"
	"go-Phones/database"
	"os"
	"strconv"
)

// Execute обрабатывает команды CLI
func Execute(db *sql.DB) error {
	args := os.Args[1:] // Пропускаем первый аргумент (имя программы)

	if len(args) == 0 {
		return errors.New("no command provided")
	}

	switch args[0] {
	case ".add-phone":
		if len(args) != 4 { // Исправлено на 4 аргумента
			return errors.New("usage: .add-phone <model> <brand> <price>")
		}

		price, err := strconv.Atoi(args[3]) // Исправлено на args[3]
		if err != nil {
			return fmt.Errorf("invalid price: %v", err)
		}

		phone := database.Phone{
			Model: args[1], // Исправлено на args[1]
			Brand: args[2], // Исправлено на args[2]
			Price: price,
		}

		if err := database.AddPhone(db, phone); err != nil {
			return fmt.Errorf("failed to add phone: %v", err)
		}

		fmt.Println("Phone added successfully!")
		return nil

	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}
