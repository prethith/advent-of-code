package days

import "advent_of_code/solution"

func GetSolution(day int) solution.Solution {
	switch day {
		case 1: return Day01{}
		default: return nil
	}
}
