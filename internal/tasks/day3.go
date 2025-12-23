package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Day3Task1 struct{}

type Bank struct {
	batteries []int
}

func (b Bank) Joltage() (int, error) {
	if len(b.batteries) < 2 {
		return -1, fmt.Errorf("invalid bank (expected at least 2 batteries): %s", b)
	}
	firstIdx, firstVal, err := max(b.batteries[:len(b.batteries)-1])
	if err != nil {
		return -1, fmt.Errorf("can't find first value, error: %s", err)
	}
	_, lastVal, err := max(b.batteries[firstIdx+1:])
	if err != nil {
		return -1, fmt.Errorf("can't find last value, error: %s", err)
	}
	return firstVal*10 + lastVal, nil
}

func (b Bank) String() string {
	sb := strings.Builder{}
	sb.WriteString("Bank{")
	for _, battery := range b.batteries {
		sb.WriteRune(rune('0' + battery))
	}
	sb.WriteString("}")
	return sb.String()
}

func BankFromString(s string) (*Bank, error) {
	b := Bank{batteries: make([]int, len(s))}
	for idx, r := range s {
		i := int(r - '0')
		if i < 0 || i > 9 {
			return nil, fmt.Errorf("expected single digit number for %c", r)
		}
		b.batteries[idx] = i
	}
	return &b, nil
}

func max(s []int) (int, int, error) {
	if len(s) == 0 {
		return -1, -1, fmt.Errorf("expected slice of positive length")
	}
	maxVal := s[0]
	maxIdx := 0
	for idx, i := range s {
		if i > maxVal {
			maxVal = i
			maxIdx = idx
		}
	}
	return maxIdx, maxVal, nil
}

func (*Day3Task1) GetName() string {
	return "day 3 task 1"
}

func (*Day3Task1) Run() int {
	path := "assets/personal-inputs/day3.txt"
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeCharDevice)
	if err != nil {
		panic(fmt.Sprintf("can't open %s, error: %s", path, err))
	}

	defer f.Close()

	banks := make([]*Bank, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if b, err := BankFromString(line); err != nil {
			panic(fmt.Sprintf("can't parse bank from %s, error: %s", line, err))
		} else {
			banks = append(banks, b)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("can't scan %s, error: %s", path, err))
	}

	acc := 0
	for _, b := range banks {
		if joltage, err := b.Joltage(); err != nil {
			panic(fmt.Sprintf("can't calculate joltage for %s, error: %s", b, err))
		} else {
			acc += joltage
		}
	}

	return acc
}
