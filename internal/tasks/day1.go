package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Day1Task1 struct{}

func (*Day1Task1) GetName() string {
	return "day 1 task 1"
}

func (*Day1Task1) Run() int {
	path := "assets/personal-inputs/day1/task1.txt"
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("can't open %s, error: %s", path, err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var turns []int
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[:1]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(fmt.Sprintf("invalid input: %s, can't convert %s to int", line, line[1:]))
		}
		switch dir {
		case "L":
			turns = append(turns, -num)
		case "R":
			turns = append(turns, num)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("can't scan %s, error: %s", path, err))
	}

	curr := 50
	acc := 0
	for _, t := range turns {
		curr = (curr + t) % 100
		if curr == 0 {
			acc++
		}
	}

	return acc
}
