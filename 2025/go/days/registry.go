package days

type Solution struct {
	Part1 func(string) (string, error)
	Part2 func(string) (string, error)
}

var solutions = map[int]Solution{
	1: {day01Part1, day01Part2},
	2: {day02Part1, day02Part2},
	3: {day03Part1, day03Part2},
	4: {day04Part1, day04Part2},
	5: {day05Part1, day05Part2},
	6: {day06Part1, day06Part2},
	7: {day07Part1, day07Part2},
}

func Get(day int) (Solution, bool) {
	s, ok := solutions[day]
	return s, ok
}
