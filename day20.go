package adventofcode2016

import (
	"slices"
)

type _range struct {
	lower, upper uint32
}

func Day20(lines []string, part1 bool) (uint32, error) {
	// Pre-allocate ranges slice
	ranges := make([]_range, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		// Parse "lower-upper" inline
		var lower, upper uint32
		i := 0
		for i < len(line) && line[i] != '-' {
			lower = lower*10 + uint32(line[i]-'0')
			i++
		}
		i++ // skip '-'
		for i < len(line) {
			upper = upper*10 + uint32(line[i]-'0')
			i++
		}
		ranges = append(ranges, _range{lower, upper})
	}

	slices.SortFunc(ranges, func(a, b _range) int {
		if a.lower != b.lower {
			return int(a.lower) - int(b.lower)
		}
		return int(a.upper) - int(b.upper)
	})

	var end, gaps, first uint32
	for _, r := range ranges {
		if end < r.lower {
			if gaps == 0 {
				first = r.lower - 1
			}
			gaps += r.lower - end
		}
		end = max(end, r.upper+1)
	}
	if part1 {
		return first, nil
	}
	return gaps, nil
}
