package days

import (
	"fmt"
	"strconv"
	"strings"
)

func day06Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// last line is operators
	ops := strings.Fields(lines[len(lines)-1])
	numProblems := len(ops)

	// collect all numbers in the columns
	problems := make([][]int, numProblems)
	for _, line := range lines[:len(lines)-1] {
		tokens := strings.Fields(line)
		for i, tok := range tokens {
			n, _ := strconv.Atoi(tok)
			problems[i] = append(problems[i], n)
		}
	}

	grandTotal := 0
	for i, nums := range problems {
		result := nums[0]
		op := ops[i]
		for _, n := range nums[1:] {
			if op == "*" {
				result *= n
			} else {
				result += n
			}
		}
		grandTotal += result
	}

	return strconv.Itoa(grandTotal), nil
}

func day06Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	ops := strings.Fields(lines[len(lines)-1])
	numberRows := lines[:len(lines)-1]

	// find the max width
	width := 0
	for _, r := range numberRows {
		if len(r) > width {
			width = len(r)
		}
	}

	// transpose for each character, by column
	transposed := make([]string, width)
	for col := 0; col < width; col++ {
		var sb strings.Builder
		for _, row := range numberRows {
			if col < len(row) {
				sb.WriteByte(row[col])
			} else {
				sb.WriteByte(' ') // if the current row falls short of the max width, add spaces
			}
		}
		transposed[col] = sb.String()
	}

	// transposed looks something like
	// transposed[0] = "1__"
	// transposed[1] = "24_"
	// and so on.

	var problems [][]int
	var current []int

	for _, row := range transposed {
		s := strings.ReplaceAll(row, " ", "")
		if s == "" {
			if len(current) > 0 {
				problems = append(problems, current)
				current = nil
			}
			continue
		}
		n, _ := strconv.Atoi(s)
		current = append(current, n)
	}

	if len(current) > 0 {
		problems = append(problems, current)
	}

	total := 0

	for i, nums := range problems {
		var result int
		switch ops[i] {
		case "+":
			result = 0
			for _, n := range nums {
				result += n
			}
		case "*":
			result = 1
			for _, n := range nums {
				result *= n
			}
		default:
			return "", fmt.Errorf("Unknown operator %q", ops[i])
		}
		total += result
	}
	return strconv.Itoa(total), nil
}
