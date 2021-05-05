package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show the task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Task List")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
