package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)

		if id == -1 {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}
		if err != nil {
			fmt.Println("Task not stored. Error:", err)
			os.Exit(1)
		}

		fmt.Printf("Task added \"%s\" to the task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
