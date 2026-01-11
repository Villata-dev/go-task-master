package cmd

import (
	"fmt"
	"go-task-master/internal/db"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Marks a task as completed.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID: %w", err)
		}

		if err := db.CompleteTaskDB(id); err != nil {
			return err
		}

		fmt.Println("Task marked as completed.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
