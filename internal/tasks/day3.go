package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Day3Task1 struct{}
type Day3Task2 struct{}

type Bank struct {
	batteries []int
}

func (b Bank) Joltage() (int, error) {
	if len(b.batteries) < 2 {
		return 0, fmt.Errorf("invalid bank (expected at least 2 batteries): %s", b)
	}
	firstIdx, firstVal := maxIntSlice(b.batteries[:len(b.batteries)-1])
	_, lastVal := maxIntSlice(b.batteries[firstIdx+1:])
	return firstVal*10 + lastVal, nil
}

func helper(batteries []int, digits int) ([]int, error) {
	if len(batteries) < digits {
		return nil, fmt.Errorf("can't calculate %d digit joltage", digits)
	}
	if digits == 0 {
		return nil, fmt.Errorf("unexpected state")
	}
	idxForNth, valForNth := maxIntSlice(batteries[:len(batteries)-digits+1])
	if digits == 1 {
		return []int{valForNth}, nil
	}
	rest, err := helper(batteries[idxForNth+1:], digits-1)
	if err != nil {
		return nil, err
	}
	return append(rest, valForNth), nil
}

func (b Bank) LargeJoltage() (int, error) {
	if len(b.batteries) < 12 {
		return 0, fmt.Errorf("invalid bank (expected at least 12 batteries for large joltage): %s", b)
	}
	joltage, err := helper(b.batteries, 12)
	if err != nil {
		return 0, err
	}
	acc := 0
	for i := len(joltage) - 1; i >= 0; i-- {
		acc = acc*10 + joltage[i]
	}
	return acc, nil
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

func BankFromString(s string) (Bank, error) {
	b := Bank{batteries: make([]int, len(s))}
	for i := 0; i < len(s); i++ {
		n := int(s[i]) - int('0')
		if n < 0 || n > 9 {
			return Bank{}, fmt.Errorf("expected single digit number for %c", s[i])
		}
		b.batteries[i] = n
	}
	return b, nil
}

func maxIntSlice(s []int) (int, int) {
	maxVal := s[0]
	maxIdx := 0
	for idx, i := range s {
		if i > maxVal {
			maxVal = i
			maxIdx = idx
		}
	}
	return maxIdx, maxVal
}

func readBanks() ([]Bank, error) {
	path := "assets/personal-inputs/day3.txt"
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't open %s, error: %w", path, err)
	}

	defer f.Close()

	var banks []Bank

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		b, err := BankFromString(line)
		if err != nil {
			return nil, fmt.Errorf("can't parse bank from %s, error: %w", line, err)
		}
		banks = append(banks, b)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("can't scan %s, error: %w", path, err)
	}

	return banks, nil
}

func (*Day3Task1) GetName() string {
	return "day 3 task 1"
}

func (*Day3Task1) Run() (int, error) {
	banks, err := readBanks()
	if err != nil {
		return 0, fmt.Errorf("failed to read banks: %w", err)
	}

	acc := 0
	for _, b := range banks {
		joltage, err := b.Joltage()
		if err != nil {
			return 0, fmt.Errorf("can't calculate joltage for %s, error: %w", b, err)
		}
		acc += joltage
	}

	return acc, nil
}

func (*Day3Task2) GetName() string {
	return "day 3 task 2"
}

func (*Day3Task2) Run() (int, error) {
	banks, err := readBanks()
	if err != nil {
		return 0, fmt.Errorf("failed to read banks")
	}

	acc := 0
	for _, b := range banks {
		joltage, err := b.LargeJoltage()
		if err != nil {
			return 0, fmt.Errorf("can't calculate joltage for %s, error: %w", b, err)
		}
		acc += joltage
	}

	return acc, nil
}
