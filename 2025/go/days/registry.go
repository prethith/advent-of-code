package days

import "advent_of_code/solution"

func GetSolution(day int) solution.Solution {
	switch day {
		case 1: return Day01{}
		case 2: return Day02{}
		case 3: return Day03{}
		default: return nil
	}
}
