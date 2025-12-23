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

func (r *Range) MergeInclusive(o *Range) (bool, *Range) {
	var smaller *Range
	var larger *Range
	if r.b < o.b {
		smaller = r
		larger = o
	} else {
		smaller = o
		larger = r
	}
	if smaller.e >= larger.b {
		return true, &Range{minInt(smaller.b, larger.b), maxInt(smaller.e, larger.e)}
	} else {
		return false, nil
	}
}

func (r Range) LenInclusive() int {
	return r.e - r.b + 1
}

func (*Day5Task1) Run() (int, error) {
	ranges, ids, _ := readInput()

	acc := 0
	for _, i := range ids {
		for _, r := range ranges {
			if r.ContainsInclusive(i) {
				acc++
				break
			}
		}
	}

	return acc, nil
}

type Day5Task2 struct{}

func (*Day5Task2) GetName() string {
	return "day 5 task 2"
}

func (*Day5Task2) Run() (int, error) {
	ranges, _, _ := readInput()
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].b < ranges[j].b
	})

	mergedRanges := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		isMergeAble, merged := ranges[i].MergeInclusive(&mergedRanges[len(mergedRanges)-1])
		if !isMergeAble {
			mergedRanges = append(mergedRanges, ranges[i])
		} else {
			mergedRanges[len(mergedRanges)-1] = *merged
		}
	}

	acc := 0
	for _, r := range mergedRanges {
		acc += r.LenInclusive()
	}

	return acc, nil
}

func readInput() ([]Range, []int, error) {
	path := "assets/personal-inputs/day5.txt"
	input, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("can't read %s, error: %w", path, err)
	}

	rawParts := bytes.Split(input, []byte{'\n', '\n'})
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
