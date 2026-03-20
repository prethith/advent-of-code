package days

import (
	"fmt"
	"strconv"
	"strings"
)

type Day02 struct{}

func repeatingPatternSums(input string) (int, int) {
	p1sum, p2sum := 0,0
	for s := range strings.SplitSeq(strings.TrimSpace(input), ",") {
		var lo, hi int
		fmt.Sscanf(s, "%d-%d", &lo, &hi)

		for i:= lo; i <= hi; i++ {
			s := strconv.Itoa(i)
			if s[:len(s)/2] == s[len(s)/2:] {
				p1sum += i
			}

			// the logic
			// if a string is repeating, then shift and trim by 1 and the original string will still be in there
			// for example, s = "abab"
			// s+s, shifted = "bababa"
			// and s is in there
			// if s = "abcd"
			// s+s shifted = "bcdabc" -> but s isn't in here
			if strings.Contains((s+s)[1:len(s+s)-1], s) {
				p2sum += i
			}
		}
	}

	return p1sum, p2sum
}

func (d Day02) Part1(input string) (string, error) {
	sum, _ := repeatingPatternSums(input)
	return strconv.Itoa(sum), nil

}

func (d Day02) Part2(input string) (string, error) {
	_, sum := repeatingPatternSums(input)
	return strconv.Itoa(sum), nil
}
