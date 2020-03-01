package cmd

import (
	"fmt"
	"strconv"

	"github.com/dhanusaputra/go-exercises/task/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("failed to parse the argument:", id)
				continue
			}
			ids = append(ids, id)
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Marked \"%d\" as completed\n.", id)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
