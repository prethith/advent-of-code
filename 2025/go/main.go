package main

import (
	"advent_of_code/days"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	day, _ := strconv.Atoi(os.Args[1])

	test := len(os.Args) > 2 && os.Args[2] == "--test"

	suffix := ""
	if test {
		suffix = "_test"
	}

	inputPath := fmt.Sprintf("../inputs/day%02d%s.txt", day, suffix)
	inputBytes, err := os.ReadFile(inputPath)

	if err != nil {
		fmt.Printf("Could not read %s\n", inputPath)
		os.Exit(1)
	}

	input := string(inputBytes)

	solution := days.GetSolution(day)
	if solution == nil {
		fmt.Printf("Day %d not implemented yet.", day)
		return
	}

	fmt.Printf("**=========================Day %d=========================**\n", day)
	start := time.Now()
	part1, err := solution.Part1(input)
	elapsed := time.Since(start)
	if err != nil {
	    fmt.Printf("  Part 1 error: %s\n", err)
	} else {
	    fmt.Printf("  Part 1: %s\n", part1)
		fmt.Printf("  Time taken: %v\n", elapsed)
	}

	start = time.Now()
	part2, err := solution.Part2(input)
	elapsed = time.Since(start)
	if err != nil {
	    fmt.Printf("  Part 2 error: %s\n", err)
	} else {
	    fmt.Printf("\n  Part 2: %s\n", part2)
		fmt.Printf("  Time taken: %v\n", elapsed)
	}
}
