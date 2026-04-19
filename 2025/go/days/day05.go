package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day05 struct{}

type Interval struct {
	start int
	end   int
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i,j int) bool {
		return intervals[i].start <= intervals[j].start
	})

	merged := []Interval{intervals[0]}

	for _, curr := range intervals[1:] {
		last := &merged[len(merged)-1]

		// no overlap
		if curr.start > last.end {
			merged = append(merged, curr)
		} else {
			// overlap
			if curr.end > last.end {
				last.end = curr.end
			}
		}
	}

	return merged
}

func (d Day05) Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var intervals []Interval
	var ids []int
	//---------------------
	// Parsing intervals
	//---------------------
	parsingIntervals := true

	for _, line := range lines {
		line := strings.TrimSpace(line)

		if line == "" {
			parsingIntervals = false
			continue
		}

		if parsingIntervals {
			parts := strings.Split(line, "-")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			intervals = append(intervals, Interval{a,b})
		} else {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}

	// merge
	merged := mergeIntervals(intervals)

	// count
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

func (d Day05) Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var intervals []Interval
	//---------------------
	// Parsing intervals
	//---------------------

	for _, line := range lines {
		line := strings.TrimSpace(line)

		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		intervals = append(intervals, Interval{a,b})
	}

	// merge
	merged := mergeIntervals(intervals)
	count := 0

	for _, span := range merged{
		count += (span.end - span.start + 1)
	}

	return fmt.Sprint(count), nil
}
