package tasks

import (
	"bytes"
	"fmt"
	"os"
)

const (
	Roll  byte = '@'
	Empty byte = '.'
)

type Day4Task1 struct{}

func (*Day4Task1) GetName() string {
	return "day 4 task 1"
}

func (*Day4Task1) Run() (int, error) {
	m, err := readMap()
	if err != nil {
		return 0, fmt.Errorf("can't parse map: %w", err)
	}
	acc := 0
	for y, row := range m {
		for x, cell := range row {
			if cell == Empty {
				continue
			}
			_acc := 0
			for j := _max(0, y-1); j < _min(len(m), y+2); j++ {
				for i := _max(0, x-1); i < _min(len(row), x+2); i++ {
					if i == x && j == y {
						continue
					}
					if m[j][i] == Roll {
						_acc++
					}
				}
			}
			if _acc < 4 {
				acc++
			}
		}
	}
	return acc, nil
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func readMap() ([][]byte, error) {
	path := "assets/personal-inputs/day4.txt"
	input, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s, error: %w", path, err)
	}
	lines := bytes.Split(input, []byte{'\n'})
	return lines, nil
}
