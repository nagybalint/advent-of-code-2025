package tasks

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day2Task1 struct{}

func (*Day2Task1) GetName() string {
	return "day 2 task 1"
}

type Range struct {
	b, e int
}

func RangeFromString(s string) (*Range, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid range: %s", s)
	}

	b, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid lower bound: %s", parts[0])
	}

	e, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid upper bound: %s", parts[1])
	}

	return &Range{b, e}, nil
}

func (r Range) String() string {
	return fmt.Sprintf("Range{b: %d, e: %d}", r.b, r.e)
}

func (r *Range) CountInvalid() int {
	acc := 0
	for i := r.b; i <= r.e; i++ {
		s := strconv.Itoa(i)
		l := len(s)
		if l^2 == 1 {
			// even digits, must be valid
			continue
		}
		start := s[:l/2]
		end := s[l/2:]
		if start == end {
			acc += i
		}
	}
	return acc
}

func (*Day2Task1) Run() int {
	path := "assets/personal-inputs/day2.txt"
	b, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("can't open %s, error: %s", path, err))
	}

	s := string(b)
	rawRanges := strings.Split(strings.TrimSpace(s), ",")

	acc := 0
	for _, raw := range rawRanges {
		r, err := RangeFromString(raw)
		if err != nil {
			panic(fmt.Sprintf("can't parse range from %s, error: %s", raw, err))
		}
		acc += r.CountInvalid()
	}

	return acc
}
