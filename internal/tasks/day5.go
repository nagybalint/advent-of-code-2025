package tasks

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Day5Task1 struct{}

func (*Day5Task1) GetName() string {
	return "day 5 task 1"
}

func (r Range) ContainsInclusive(x int) bool {
	if x >= r.b && x <= r.e {
		return true
	}
	return false
}

func (*Day5Task1) Run() (int, error) {
	ranges, ids, _ := readInput()
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].b > ranges[j].b
	})

	acc := 0 
	for _, i := range(ids) {
		for _, r := range(ranges) {
			if r.ContainsInclusive(i) {
				acc++
				break
			}
		}
	}

	return acc, nil
}

func readInput() ([]Range, []int, error) {
	path := "assets/personal-inputs/day5.txt"
	input, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("can't read %s, error: %w", path, err)
	}

	rawParts := bytes.Split(input, []byte{'\n','\n'})
	rawRanges := bytes.Split(rawParts[0], []byte{'\n'})
	rawIds := bytes.Split(rawParts[1], []byte{'\n'})

	var ranges []Range
	for _, rawRange := range rawRanges {
		r, err := RangeFromString(string(rawRange))
		if err != nil {
			return nil, nil, fmt.Errorf("can't parse Range from %s, error: %w", rawRange, err)
		}
		ranges = append(ranges, *r)
	}

	var ids []int
	for _, rawId := range rawIds {
		id, err := strconv.Atoi(string(rawId))
		if err != nil {
			return nil, nil, fmt.Errorf("can't parse id from %s, error: %w", rawId, err)
		}
		ids = append(ids, id)
	}
	return ranges, ids, nil
}
