package commands

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Root",
	Short: "A CLI tool for managing phones database",
	Long:  `This is a CLI tool for managing the phones database using SQLite.`,
}

func Execute() error {
	return RootCmd.Execute()
}
