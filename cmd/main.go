package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nagybalint/advent-of-code-2025/internal/tasks"
)

type Task interface {
	GetName() string
	Run() int32
}

var taskRegistry = map[string]map[string]func() Task{
	"1": {
		"1": func() Task { return &tasks.Day1Task1{} },
	},
}

func main() {
	var day string
	var taskName string
	cmd := cobra.Command{
		Use:   "aoc25",
		Short: "Run a single task of Advent of Code 2025",
		RunE: func(cmd *cobra.Command, args []string) error {
			taskC := taskRegistry[day][taskName]
			task := taskC()
			fmt.Printf("The result of task %s is %d \n", task.GetName(), task.Run())
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&day, "day", "d", "day selector")
	cmd.PersistentFlags().StringVar(&taskName, "task", "t", "task selector")
	cmd.Execute()
}
