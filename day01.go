package adventofcode2016

import (
	"math"
	"strconv"
	"strings"
)

func Day01(line string, part1 bool) uint {
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
		n, _ := strconv.Atoi(s)
		switch turn {
		case 'R':
			facing *= 0 - 1i
		case 'L':
			facing *= 0 + 1i
		}
		for i := 0; i < n; i++ {
			here += facing
			if !part1 {
				if positions[here] {
					return manhattanDistance()
				}
				positions[here] = true
			}
		}
	}
	return manhattanDistance()
}
