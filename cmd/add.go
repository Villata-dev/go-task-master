package cmd

import (
	"fmt"
	"go-task-master/internal/db"
	"go-task-master/internal/models"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [title] [description]",
	Short: "Adds a new task.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		task := models.Task{
			Title:       args[0],
			Description: args[1],
			Completed:   false,
		}

		if _, err := db.CreateTaskDB(task); err != nil {
			return err
		}

		fmt.Println("Task added successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
