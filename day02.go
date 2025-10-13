package adventofcode2016

// Day02 solves day 2, both part 1 and part 2.
func Day02(input []byte, part1 bool) string {
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

	var code []byte
	for _, b := range input {
		if b == '\n' {
			// End of line - record digit
			digit := digits[here]
			code = append(code, digit)
		} else if step, ok := steps[b]; ok {
			// Valid movement instruction
			if inBound(here + step) {
				here += step
			}
		}
		// Ignore invalid characters
	}
	// Handle final line if input doesn't end with newline
	if len(input) > 0 && input[len(input)-1] != '\n' {
		digit := digits[here]
		code = append(code, digit)
	}

	// Return "0" if no valid output was generated (garbage input)
	if len(code) == 0 {
		return "0"
	}
	return string(code)
}
