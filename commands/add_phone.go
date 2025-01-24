package commands

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"go-Phones/database"
)

var (
	name  string
	brand string
	price float64
)

func init() {
	addPhoneCmd.Flags().StringVarP(&name, "name", "n", "", "Phone name (required)")
	addPhoneCmd.Flags().StringVarP(&brand, "brand", "b", "", "Phone brand (required)")
	addPhoneCmd.Flags().Float64VarP(&price, "price", "p", 0.0, "Phone price (required)")

	err := addPhoneCmd.MarkFlagRequired("name")
	if err != nil {
		return
	}
	err2 := addPhoneCmd.MarkFlagRequired("brand")
	if err2 != nil {
		return
	}
	err3 := addPhoneCmd.MarkFlagRequired("price")
	if err3 != nil {
		return
	}

	RootCmd.AddCommand(addPhoneCmd)
}

var addPhoneCmd = &cobra.Command{
	Use:   "add-phone",
	Short: "Add a new phone to the database",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.ConnectDB()
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			return
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)

		_, err = db.Exec("INSERT INTO phones (name, brand, price) VALUES (?, ?, ?)", name, brand, price)
		if err != nil {
			fmt.Printf("Error inserting phone into database: %v\n", err)
			return
		}

		fmt.Println("Phone added successfully!")
	},
}
