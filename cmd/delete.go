package cmd

import (
	"fmt"
	"go-task-master/internal/db"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Deletes a task by its ID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID: %w", err)
		}

		if err := db.DeleteTaskDB(id); err != nil {
			return err
		}

		fmt.Println("Task deleted successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
