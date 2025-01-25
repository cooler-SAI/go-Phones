package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"go-Phones/shlex"
	"os"
)

func Commands(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter commands (e.g., .add-phone <model> <brand> <price>):")
	for scanner.Scan() {
		input := scanner.Text()

		args, err := shlex.Split(input)
		if err != nil {
			fmt.Println("Failed to parse input:", err)
			continue
		}

		if err := shlex.ProcessCommand(db, args); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
