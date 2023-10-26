package adventofcode2016

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Trying to find a solution in the bitmask space, to no avail...

type _range struct {
	lower, upper uint32
}

func Day20(lines []string, part1 bool) (uint32, error) {
	var ranges []_range
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		parts := strings.Split(lines[i], "-")
		if len(parts) != 2 {
			return 0, fmt.Errorf("line %d: want two ranges but got %q", i, lines[i])
		}

		a, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, fmt.Errorf("line %d: first operand not a number: %q", i, lines[i])
		}
		b, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, fmt.Errorf("line %d: first operand not a number: %q", i, lines[i])
		}
		ranges = append(ranges, _range{uint32(a), uint32(b)})
	}

	sort.Slice(ranges, func(i, j int) bool {
		// two key sort
		if ranges[i].lower == ranges[j].lower {
			return ranges[i].upper < ranges[j].upper
		}
		return ranges[i].lower < ranges[j].lower
	})

	var end, gaps, first uint32
	for i := range ranges {
		if end < ranges[i].lower {
			if gaps == 0 {
				first = ranges[i].lower - 1
			}
			gaps += ranges[i].lower - end
		}
		end = max(end, ranges[i].upper+1)
	}
	if part1 {
		return first, nil
	}
	return gaps, nil
}
