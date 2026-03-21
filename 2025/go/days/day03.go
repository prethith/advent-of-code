package days

import (
	"fmt"
	"strconv"
	"strings"
)

type Day03 struct{}

func (d Day03) Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0
	for _, line := range lines {
		max1, max2 := -1, -1
		for i := 0; i < len(line)-1; i++ {
			if max1 == -1 || line[i] > line[max1] {
				max1 = i
			}
		}

		for j := max1 + 1; j < len(line); j++ {
			if max2 == -1 || line[j] > line[max2] {
				max2 = j
			}
		}
		result := int((line[max1]-'0')*10 + (line[max2] - 48))
		sum += result
	}
	return strconv.Itoa(sum), nil
}

func maxNumber(s string, k int) string {
	stack := []byte{}
	// the logic
	// we use a stack
	// pop the top element if the next element is greater
	// and append the greater element
	// but only a limited number of such operations are allowed
	for i:= 0; i < len(s); i++ {
		for k > 0 && len(stack) > 0 && s[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, s[i])
	}

	// if any removals are still left, remove from the end
	stack = stack[:len(stack)-k]

	return string(stack)
}


func (d Day03) Part2(input string) (string, error) {
	inputs := strings.Split(strings.TrimSpace(input), "\n")
	var sum int64 = 0
	for _, s := range inputs {
		var num int64
		res := maxNumber(s, (len(s)-12))

		fmt.Sscan(res, &num)
		sum += num

	}
	return strconv.FormatInt(sum, 10), nil
}
