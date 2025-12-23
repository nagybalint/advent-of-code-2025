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

type Day4Map [][]byte

type Day4Task1 struct{}

func (*Day4Task1) GetName() string {
	return "day 4 task 1"
}

func (m Day4Map) h() int {
	return len(m)
}

func (m Day4Map) w() int {
	return len(m[0])
}

func (m Day4Map) IsRemovable(x, y int) bool {
	if m[y][x] == Empty {
		return false
	}
	acc := 0
	for j := maxInt(0, y-1); j < minInt(m.h(), y+2); j++ {
		for i := maxInt(0, x-1); i < minInt(m.w(), x+2); i++ {
			if i == x && j == y {
				continue
			}
			if m[j][i] == Roll {
				acc++
			}
		}
	}
	if acc < 4 {
		return true
	} else {
		return false
	}
}

func (*Day4Task1) Run() (int, error) {
	m, err := readMap()
	if err != nil {
		return 0, fmt.Errorf("can't parse map: %w", err)
	}
	acc := 0
	for y, row := range m {
		for x := range row {
			if m.IsRemovable(x, y) {
				acc++
			}
		}
	}
	return acc, nil
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func readMap() (Day4Map, error) {
	path := "assets/personal-inputs/day4.txt"
	input, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s, error: %w", path, err)
	}
	lines := bytes.Split(input, []byte{'\n'})
	return Day4Map(lines), nil
}

type Day4Task2 struct{}

func (*Day4Task2) GetName() string {
	return "day 4 task 2"
}

func (*Day4Task2) Run() (int, error) {
	m, err := readMap()
	if err != nil {
		return 0, fmt.Errorf("can't parse map: %w", err)
	}
	acc := 0
	for {
		var removables [][2]int
		for y, row := range m {
			for x := range row {
				if m.IsRemovable(x, y) {
					removables = append(removables, [2]int{x, y})
				}
			}
		}
		if len(removables) > 0 {
			acc += len(removables)
			for _, pair := range removables {
				x := pair[0]
				y := pair[1]
				m[y][x] = Empty
			}
		} else {
			break
		}
	}
	return acc, nil
}
