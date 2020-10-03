package adventofcode2016

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Day1Part1 solves day 1 part 1.
func Day1Part1(line string) (uint, error) {
	return day1(line, false)
}

// Day1Part2 solves day 1 part 1.
func Day1Part2(line string) (uint, error) {
	return day1(line, true)
}

func day1(line string, part2 bool) (uint, error) {
	here := 0 + 0i
	abs := func(f float64) uint {
		return uint(math.Abs(f))
	}
	manhattanDistance := func() uint {
		return abs(real(here)) + abs(imag(here))
	}
	facing := 0 + 1i
	positions := make(map[complex128]bool)
	for _, s := range strings.Split(line, ", ") {
		buf := []byte(s)
		turn := buf[0]
		s := string(buf[1:])
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		switch turn {
		case 'R':
			facing *= 0 - 1i
		case 'L':
			facing *= 0 + 1i
		default:
			return 0, fmt.Errorf("unknown turn %q", turn)
		}
		// do not mark end position only, but all interim steps
		for i := 0; i < n; i++ {
			here += facing
			if part2 {
				if positions[here] {
					return manhattanDistance(), nil
				}
				positions[here] = true
			}
		}
	}
	return manhattanDistance(), nil
}
