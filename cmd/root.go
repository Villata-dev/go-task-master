package cmd

import (
	"fmt"
	"go-task-master/internal/db"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-master",
	Short: "Task Master is a simple task management CLI.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "serve" || cmd.Name() == "help" || cmd.Name() == "completion" {
			return nil
		}
		return db.InitDB("./tasks.db")
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Default to serve if no command is specified
		serveCmd.Run(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
