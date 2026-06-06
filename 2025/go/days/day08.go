package days

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var NUM_ITERATIONS_SAMPLE int = 10
var NUM_ITERATIONS_MAIN int = 1000

type Point struct {
	X int
	Y int
	Z int
}

// New returns a Point based on the parsed input
func New(x, y, z string) Point {
	X, _ := strconv.Atoi(x)
	Y, _ := strconv.Atoi(y)
	Z, _ := strconv.Atoi(z)
	return Point{X: X, Y: Y, Z: Z}
}

func (p Point) Distance(p2 Point) float64 {
	dx := float64(p2.X - p.X)
	dy := float64(p2.Y - p.Y)
	dz := float64(p2.Z - p.Z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Union Find
// DSU is Disjoint Set Union, another name for Union Find
type DSU struct {
	parent []int
	size   []int
	count  int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	return &DSU{
		parent: parent,
		size:   size,
		count:  n,
	}
}

func (d *DSU) Find(v int) int {
	if d.parent[v] != v {
		d.parent[v] = d.Find(d.parent[v])
	}
	return d.parent[v]
}

func (d *DSU) Union(a, b int) bool {
	a = d.Find(a)
	b = d.Find(b)
	if a == b {
		return false
	}

	if d.size[a] < d.size[b] {
		a, b = b, a
	}
	d.parent[b] = a
	d.size[a] += d.size[b]
	d.count --
	return true
}

// to precompute distances
type Edge struct {
	A int
	B int
	D float64
}

func parseInput(input string) ([]Point, []Edge) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		coords := strings.Split(strings.TrimSpace(line), ",")
		points = append(points, New(coords[0], coords[1], coords[2]))
	}

	edges := make([]Edge, 0, len(points)*(len(points)-1)/2)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				A: i,
				B: j,
				D: points[i].Distance(points[j]),
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].D < edges[j].D
	})

	return points, edges
}

func day08Part1(input string) (string, error) {
	points, edges := parseInput(input)

	dsu := NewDSU(len(points))

	for i := 0; i < NUM_ITERATIONS_MAIN; i++ {
		dsu.Union(edges[i].A, edges[i].B)
	}

	sizes := make([]int, 0)
	for i := range points {
		if dsu.Find(i) == i {
			sizes = append(sizes, dsu.size[i])
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return strconv.Itoa(sizes[0] * sizes[1] * sizes[2]), nil
}


func day08Part2(input string) (string, error) {
	points, edges := parseInput(input)

	dsu := NewDSU(len(points))

	for _, e := range edges {
		if dsu.Union(e.A, e.B) && dsu.count == 1 {
			answer := points[e.A].X * points[e.B].X
			return strconv.Itoa(answer), nil
		}
	}
	return "", fmt.Errorf("The graph never became connected")
}
