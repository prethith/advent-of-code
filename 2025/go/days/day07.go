package days

import (
	"strconv"
	"strings"
)

func day07Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sLoc := strings.Index(lines[0], "S")

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	indices := []int{sLoc}
	splits := 0

	for row := 1; row < len(grid); row++ {
		next := make([]int, 0)

		for _, col := range indices {
			if col < 0 || col >= len(grid[row]) {
				continue
			}

			switch grid[row][col] {
			case '.':
				grid[row][col] = '|'
				next = append(next, col)

			case '^':
				splits++

				// visualization only
				if col > 0 {
					grid[row][col-1] = '|'
					next = append(next, col-1)
				}

				if col+1 < len(grid[row]) {
					grid[row][col+1] = '|'
					next = append(next, col+1)
				}
			}
		}

		// merge beams landing at the same column
		seen := make(map[int]bool)
		indices = indices[:0]

		for _, col := range next {
			if seen[col] {
				continue
			}

			seen[col] = true
			indices = append(indices, col)
		}
	}

	return strconv.Itoa(splits), nil
}
func day07Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sLoc := strings.Index(lines[0], "S")

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	// counts[col] --> no of timelines that have particle on this column
	counts := map[int]uint64{sLoc: 1}
	for row := 1; row < len(grid); row++ {
		next := make(map[int]uint64)
		add := func(col int, n uint64) {
			if col < 0 || col >= len(grid[row]) {
				return
			}
			next[col] += n
		}

		for col, n := range counts {
			if col < 0 || col >= len(grid[row]) {
				continue
			}
			switch grid[row][col] {
			case '^':
				add(col-1, n)
				add(col+1, n)
			default:
				add(col, n)
			}
		}
		counts = next
	}
	var total uint64
	for _, n := range counts {
		total += n
	}

	return strconv.FormatUint(total, 10), nil
}
