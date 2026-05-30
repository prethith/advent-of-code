package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	start int
	end   int
}

func mergeIntervals(intervals []interval) []interval {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start <= intervals[j].start
	})

	merged := []interval{intervals[0]}
	for _, curr := range intervals[1:] {
		last := &merged[len(merged)-1]
		if curr.start > last.end {
			merged = append(merged, curr)
		} else if curr.end > last.end {
			last.end = curr.end
		}
	}

	return merged
}

func parseIntervals(lines []string) []interval {
	var intervals []interval
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		intervals = append(intervals, interval{a, b})
	}
	return intervals
}

func day05Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var ids []int
	parsingIntervals := true
	var intervalLines []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			parsingIntervals = false
			continue
		}
		if parsingIntervals {
			intervalLines = append(intervalLines, line)
		} else {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}

	merged := mergeIntervals(parseIntervals(intervalLines))

	count := 0
	for _, id := range ids {
		for _, in := range merged {
			if id >= in.start && id <= in.end {
				count++
				break
			}
		}
	}

	return fmt.Sprint(count), nil
}

func day05Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	merged := mergeIntervals(parseIntervals(lines))

	count := 0
	for _, span := range merged {
		count += span.end - span.start + 1
	}

	return fmt.Sprint(count), nil
}
