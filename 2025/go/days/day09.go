package days

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sort2(a,b int) (lo, hi int) {
	if a > b {
		return b,a
	}
	return a,b
}

// returns sorted distinct values
func sortedUnique(get func(i int) int, n int) []int {
	seen := make(map[int]struct{}, n)
	for i := range n {
		seen[get(i)] = struct{}{}
	}
	out := make([]int, 0, len(seen))
	for v := range seen {
		out = append(out, v)
	}
	sort.Ints(out)
	return out
}

// returns a look up table --> mapping from real coordinate to compressed coordinate
// Given [2,7,9,11] -> {2:0, 7:1, 9:2, 11:3}
func indexOf(vals []int) map[int]int {
	m := make(map[int]int, len(vals))
	for i, v := range vals {
		m[v] = i
	}
	return m
}

func day09Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		coords := strings.Split(strings.TrimSpace(line), ",")
		points = append(points, New(coords[0], coords[1], "0")) // Z set to 0 for 2D points. check day08.go
	}

	maxArea := 0.0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			l := math.Abs(float64(points[i].X - points[j].X + 1))
			w := math.Abs(float64(points[i].Y - points[j].Y + 1))
			area := l * w
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return strconv.FormatFloat(maxArea, 'f', -1, 64), nil
}

func day09Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	points := make([]Point, 0, len(lines)) // each line is a "x,y" point
	for _, line := range lines {
		coords := strings.Split(strings.TrimSpace(line), ",")
		points = append(points, New(coords[0], coords[1], "0")) // Z set to 0 for 2D points. check day08.go
	}
	n := len(points)

	// COORDINATE COMPRESSION
	// We create a grid of all the points, but instead of using the whole set of coordinates,
	// we compress it down by keeping only the ones that matter.
	// we only store the actual points -- in their x and y values,
	// while we assume that there is space between them that represents the other points

	xs := sortedUnique(func(i int) int { return points[i].X }, n)
	ys := sortedUnique(func(i int) int { return points[i].Y }, n)
	xi := indexOf(xs)  // real X -> compressed column mapping
	yi := indexOf(ys) // real Y -> compressed row mapping

	width, height := len(xs), len(ys)
	grid := make([][]byte, width)
	for i := range grid {
		grid[i] = make([]byte, height)
	}

	// fill in the grid with 1s where red/green
	for i := range n {
		p1, p2 := points[i], points[(i+1)%n] // every pair of points, wrapping to the front
		cxLo, cxHi := sort2(xi[p1.X], xi[p2.X]) // sort X coords
		cyLo, cyHi := sort2(yi[p1.Y], yi[p2.Y]) // sort X coords

		// eg: (7,1) -> (1,0) and (11,1) -> (3,0)
		// cxLo, cxHi = 1,3
		// cyLo, cyHi = 0,0
		// this fills (1,0), (2,0), (3,0)
		// these are just the border points
		for cx := cxLo; cx <= cxHi; cx++ {
			for cy := cyLo; cy <= cyHi; cy++ {
				grid[cx][cy] = 1
			}
		}
	}

	// flood fill for the interior points
	// we can't add a padding in compressed space
	// instead we introduce a "ring" one index outside the grid
	// on every side so it can wrap around.
	// this part discovers cells OUTSIDE the shape

	visited := make(map[[2]int]bool)
	inGrid := func(x,y int) bool { return x >= 0 && x < width && y >= 0 && y < height }
	start := [2]int{-1,-1}
	visited[start] = true
	queue := [][2]int{start}

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		for _,d := range [4][2]int{{1,0}, {-1,0}, {0,1}, {0,-1}} {
			nx, ny := c[0] + d[0], c[1] + d[1]
			if nx < -1 || ny < -1 || nx > width || ny > height {
				continue
			}
			// wall encountered
			if inGrid(nx, ny) && grid[nx][ny] == 1 {
				continue
			}
			if visited[[2]int{nx,ny}] {
				continue
			}
			visited[[2]int{nx,ny}] = true
			queue = append(queue, [2]int{nx,ny})
		}
	}

	// any cell that is NOT green is red.
	for x := range width {
		for y := range height {
			if !visited[[2]int{x,y}] {
				grid[x][y] = 1
			}
		}
	}

	// checks if all the points within a rectangle are green or not
	isAllGreen := func(cxLo, cxHi, cyLo, cyHi int) bool {
		for cx := cxLo; cx <= cxHi; cx++ {
			for cy := cyLo; cy <= cyHi; cy++ {
				if grid[cx][cy] == 0 {
					return false
				}
			}
		}
		return true
	}

	// pairwise search for every pair of points
	// calculate area on raw coordinates
	best := 0

	for i := range n {
		for j := range i {
			cxLo, cxHi := sort2(xi[points[i].X], xi[points[j].X])
			cyLo, cyHi := sort2(yi[points[i].Y], yi[points[j].Y])

			if !isAllGreen(cxLo, cxHi, cyLo, cyHi) {
				continue // not a valid rectangle
			}

			// area
			w := abs(points[i].X - points[j].X) + 1
			h := abs(points[i].Y - points[j].Y) + 1

			if w * h > best {
				best = w * h
			}
		}
	}
	return strconv.Itoa(best), nil
}
