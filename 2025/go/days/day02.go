package days

import (
	"fmt"
	"strconv"
	"strings"
)

func repeatingPatternSums(input string) (int, int) {
	p1sum, p2sum := 0, 0
	for s := range strings.SplitSeq(strings.TrimSpace(input), ",") {
		var lo, hi int
		fmt.Sscanf(s, "%d-%d", &lo, &hi)

		for i := lo; i <= hi; i++ {
			s := strconv.Itoa(i)
			if s[:len(s)/2] == s[len(s)/2:] {
				p1sum += i
			}

			// if a string is repeating, shift-and-trim by 1 still contains the original
			if strings.Contains((s+s)[1:len(s+s)-1], s) {
				p2sum += i
			}
		}
	}

	return p1sum, p2sum
}

func day02Part1(input string) (string, error) {
	sum, _ := repeatingPatternSums(input)
	return strconv.Itoa(sum), nil
}

func day02Part2(input string) (string, error) {
	_, sum := repeatingPatternSums(input)
	return strconv.Itoa(sum), nil
}
