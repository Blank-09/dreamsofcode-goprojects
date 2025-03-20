package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Blank-09/dreamsofcode-goprojects/01-todo-list/constants"
	"github.com/Blank-09/dreamsofcode-goprojects/01-todo-list/internal/csv_handler"
	"github.com/Blank-09/dreamsofcode-goprojects/01-todo-list/internal/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:     "add <description>",
	Short:   "Add a new task",
	Long:    "This `add` command adds a new task to the task list",
	Aliases: []string{"create"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]

		csv := csv_handler.CSVHandler{}
		tasks, err := csv.ReadFromCSV(constants.CsvFilePath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		task := model.Task{
			ID:          strconv.FormatInt(int64(len(tasks)+1), 10),
			Description: description,
			CreatedAt:   time.Now().Format(time.RFC3339),
			IsCompleted: false,
		}

		tasks = append(tasks, task)

		csv.WriteToCSV(constants.CsvFilePath, tasks)

		fmt.Printf("Task added: %s\n", description)
	},
}
