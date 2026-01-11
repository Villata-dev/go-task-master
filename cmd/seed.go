package cmd

import (
	"fmt"
	"go-task-master/internal/db"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seeds the database with initial data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := db.InitDB("./tasks.db"); err != nil {
			return err
		}
		if err := db.SeedTasksDB(); err != nil {
			return err
		}
		fmt.Println("Database seeded successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
