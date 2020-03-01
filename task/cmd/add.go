package cmd

import (
	"fmt"
	"strings"

	"github.com/dhanusaputra/go-exercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("something went wrong:", err.Error())
		}
		fmt.Printf("Add \"%s\" to task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
