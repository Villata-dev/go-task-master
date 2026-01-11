package cmd

import (
	"fmt"
	"go-task-master/internal/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks.",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := db.GetTasksDB()
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		for _, task := range tasks {
			status := " "
			if task.Completed {
				status = "✔"
			}
			fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
