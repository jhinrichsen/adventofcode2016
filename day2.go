package adventofcode2016

import "strings"

// Day2 solves day 2, both part 1 and part 2.
func Day2(lines []string, part1 bool) (string, error) {
	// start at '5' for both parts
	var here complex128
	if part1 {
		here = 1 + 1i
	} else {
		here = 0 + 2i
	}

	var digits map[complex128]byte
	if part1 {
		digits = map[complex128]byte{
			0 + 0i: '1',
			1 + 0i: '2',
			2 + 0i: '3',

			0 + 1i: '4',
			1 + 1i: '5',
			2 + 1i: '6',

			0 + 2i: '7',
			1 + 2i: '8',
			2 + 2i: '9',
		}
	} else {
		digits = map[complex128]byte{
			2 + 0i: '1',

			1 + 1i: '2',
			2 + 1i: '3',
			3 + 1i: '4',

			0 + 2i: '5',
			1 + 2i: '6',
			2 + 2i: '7',
			3 + 2i: '8',
			4 + 2i: '9',

			1 + 3i: 'A',
			2 + 3i: 'B',
			3 + 3i: 'C',

			2 + 4i: 'D',
		}
	}
	inBound := func(c complex128) bool {
		_, ok := digits[c]
		return ok
	}

	var steps = map[byte]complex128{
		'U': 0 - 1i,
		'R': 1 + 0i,
		'D': 0 + 1i,
		'L': -1 + 0i,
	}

	var code strings.Builder
	for _, line := range lines {
		for _, b := range []byte(line) {
			step := steps[b]
			if inBound(here + step) {
				here += step
			}
		}
		digit := digits[here]
		code.WriteByte(digit)

	}
	return code.String(), nil
}
