package days

import (
	"fmt"
	"image"
	"maps"
	"strings"
)

var delta = []image.Point{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
}

func day04Part1(input string) (string, error) {
	lines := strings.Fields(input)
	height := len(lines)
	width := len(lines[0])

	count := 0
	for y := range height {
		for x := range width {
			if lines[y][x] != '@' {
				continue
			}
			neighbours := 0
			for _, d := range delta {
				nx, ny := x+d.X, y+d.Y
				if nx >= 0 && nx < width && ny >= 0 && ny < height && lines[ny][nx] == '@' {
					neighbours++
				}
			}
			if neighbours < 4 {
				count++
			}
		}
	}

	return fmt.Sprint(count), nil
}

func day04Part2(input string) (string, error) {
	lines := strings.Fields(input)
	grid := map[image.Point]int{}
	for y, line := range lines {
		for x, char := range line {
			if char == '@' {
				grid[image.Point{X: x, Y: y}] = 1
			}
		}
	}

	removed := 0
	for {
		next := maps.Clone(grid)
		for point := range grid {
			rolls := 0
			for _, d := range delta {
				rolls += grid[point.Add(d)]
			}
			if rolls < 4 {
				delete(next, point)
				removed++
			}
		}
		if maps.Equal(next, grid) {
			break
		}
		grid = next
	}
	return fmt.Sprint(removed), nil
}
