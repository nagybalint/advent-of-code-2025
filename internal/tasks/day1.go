package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Day1Task1 struct{}
type Day1Task2 struct{}

func (*Day1Task1) GetName() string {
	return "day 1 task 1"
}

func getTurns() ([]int, error) {
	path := "assets/personal-inputs/day1.txt"
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't open %s, error: %s", path, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var turns []int
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[:1]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid input: %s, can't convert %s to int", line, line[1:])
		}
		switch dir {
		case "L":
			turns = append(turns, -num)
		case "R":
			turns = append(turns, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("can't scan %s, error: %s", path, err)
	}
	return turns, nil
}

func (*Day1Task1) Run() int {
	turns, err := getTurns()
	if err != nil {
		panic(fmt.Sprintf("error getting turns, error: %s", err))
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

func (*Day1Task2) GetName() string {
	return "day 1 task 2"
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (*Day1Task2) Run() int {
	turns, err := getTurns()
	if err != nil {
		panic(fmt.Sprintf("error getting turns, error: %s", err))
	}
	curr := 50
	acc := 0
	for _, t := range turns {
		acc += absInt(t) / 100
		next := curr + t % 100
		if curr != 0 && next <= 0 || next > 99 {
			acc += 1
		}
		curr = (next + 100) % 100
	}

	return acc
}
