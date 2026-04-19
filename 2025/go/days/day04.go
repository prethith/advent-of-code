package days

import (
	"fmt"
	"image"
	"maps"
	"strings"
)

type Day04 struct{}

func (d Day04) Part1(input string) (string, error) {
	lines := strings.Fields(input)
	height := len(lines)
	width := len(lines[0])

	// 8-neighbourhood
	delta := []image.Point{
		{0,-1}, {1,0}, {0,1}, {-1,0},
		{-1,-1}, {1,-1}, {1,1}, {-1, 1},
	}
	count := 0

	for y := range height {
		for x := range width {
			if lines[y][x] != '@' {
				continue
			}

			neighbours := 0

			for _, d := range delta {
				n_x := x + d.X
				n_y := y + d.Y

				if n_x >= 0 && n_x < width && n_y >= 0 && n_y < height {
					if lines[n_y][n_x] == '@' {
						neighbours++
					}
				}
			}

			if neighbours < 4 {
				count++
			}
		}
	}

	return fmt.Sprint(count), nil
}

func (d Day04) Part2(input string) (string, error) {
	// use a map to store the points
	lines := strings.Fields(input)
	// 8-neighbourhood
	delta := []image.Point{
		{0,-1}, {1,0}, {0,1}, {-1,0},
		{-1,-1}, {1,-1}, {1,1}, {-1, 1},
	}

	grid := map[image.Point]int{}
	for y, line := range lines {
		for x, char := range line {
			if char == '@' {
				grid[image.Point{x,y}] = 1
			}
		}
	}

	removed := 0

	for i := 0; ; i++ {
		next := maps.Clone(grid)
		for point := range grid {
			rolls := 0
			for _, d := range delta {
				rolls += grid[point.Add(d)] // returns 1 if point+d exists. 0 otherwise
			}

			if rolls < 4 {
				delete(next, point) // remove point from the grid
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
