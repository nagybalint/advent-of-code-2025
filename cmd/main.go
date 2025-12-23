package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nagybalint/advent-of-code-2025/internal/tasks"
)

type Task interface {
	GetName() string
	Run() (int, error)
}

var taskRegistry = map[string]map[string]func() Task{
	"1": {
		"1": func() Task { return &tasks.Day1Task1{} },
		"2": func() Task { return &tasks.Day1Task2{} },
	},
	"2": {
		"1": func() Task { return &tasks.Day2Task1{} },
		"2": func() Task { return &tasks.Day2Task2{} },
	},
	"3": {
		"1": func() Task { return &tasks.Day3Task1{} },
		"2": func() Task { return &tasks.Day3Task2{} },
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
			if result, err := task.Run(); err != nil {
				panic(fmt.Sprintf("Task %s failed with %s", task.GetName(), err))
			} else {
				fmt.Printf("The result of task %s is %d \n", task.GetName(), result)
			}
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&day, "day", "1", "day selector")
	cmd.PersistentFlags().StringVar(&taskName, "task", "1", "task selector")
	cmd.Execute()
}
