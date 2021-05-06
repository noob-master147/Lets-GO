package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"task/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show the task list",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.AllTask()
		if err != nil {
			fmt.Println("Something went wrong ", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no task remaining! ")
		} else {

			fmt.Println("You have the following tasks")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Value)
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
