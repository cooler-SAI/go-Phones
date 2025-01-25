package shlex

import (
	"database/sql"
	"fmt"
	"github.com/google/shlex"
	"go-Phones/database"
	"strconv"
)

// Split разбивает строку на аргументы с учётом кавычек.
func Split(input string) ([]string, error) {
	return shlex.Split(input)
}

// ProcessCommand обрабатывает команды из консоли.
func ProcessCommand(db *sql.DB, args []string) error {
	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case ".add-phone":
		if len(args) != 4 {
			return fmt.Errorf("usage: .add-phone <model> <brand> <price>")
		}

		price, err := strconv.Atoi(args[3])
		if err != nil {
			return fmt.Errorf("invalid price: %v", err)
		}

		phone := database.Phone{
			Model: args[1],
			Brand: args[2],
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
