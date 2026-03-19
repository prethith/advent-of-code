package days

import (
	"fmt"
	"strconv"
	"strings"
)

type Day01 struct{}

func (d Day01) Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var count = 0
	var dial = 50
	for _, line := range lines {
		dir := string(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("bad line %q: %w", line, err)
		}

		switch dir {
		case "L":
			dial = (dial - num + 100) % 100
		case "R":
			dial = (dial + num) % 100
		}
		if dial == 0 {
			count++
		}
	}

	return strconv.Itoa(count), nil
}
func (d Day01) Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	count := 0
	dial := 50
	for _, line := range lines {
		dir := string(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("bad line %q: %w", line, err)
		}

		count += num / 100
		num = num % 100

		switch dir {
		case "L":
			if dial != 0 && num >= dial {
				count++
			}
			dial = (dial - num + 100) % 100
		case "R":
			if dial+num >= 100 {
				count++
			}
			dial = (dial + num) % 100
		}

	}
	return strconv.Itoa(count), nil
}
