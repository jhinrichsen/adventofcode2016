package adventofcode2016

// Day03 returns number of valid triangles.
// A valid triangle has the sum of any two sides greater than the remaining side.
func Day03(input []byte, part1 bool) uint {
	// Count numbers to pre-allocate (count newlines * 3 numbers per line)
	lines := 0
	for _, b := range input {
		if b == '\n' {
			lines++
		}
	}
	if len(input) > 0 && input[len(input)-1] != '\n' {
		lines++
	}

	// Parse all numbers from input
	nums := make([]uint, 0, lines*3)
	var n uint
	inNum := false
	for _, b := range input {
		if b >= '0' && b <= '9' {
			n = n*10 + uint(b-'0')
			inNum = true
		} else if inNum {
			nums = append(nums, n)
			n = 0
			inNum = false
		}
	}
	if inNum {
		nums = append(nums, n)
	}

	// Check triangles - sum of any two sides must be greater than the third
	isValid := func(a, b, c uint) bool {
		// Sort so a <= b <= c
		if b < a {
			a, b = b, a
		}
		if c < b {
			b, c = c, b
		}
		if b < a {
			a, b = b, a
		}
		return a+b > c
	}

	var count uint
	if part1 {
		// Part 1: triangles are on each row (consecutive triples)
		for i := 0; i+2 < len(nums); i += 3 {
			if isValid(nums[i], nums[i+1], nums[i+2]) {
				count++
			}
		}
	} else {
		// Part 2: triangles are in columns (read 3 rows, 3 columns = 3 triangles)
		for i := 0; i+8 < len(nums); i += 9 {
			// Column 0: nums[i], nums[i+3], nums[i+6]
			// Column 1: nums[i+1], nums[i+4], nums[i+7]
			// Column 2: nums[i+2], nums[i+5], nums[i+8]
			if isValid(nums[i], nums[i+3], nums[i+6]) {
				count++
			}
			if isValid(nums[i+1], nums[i+4], nums[i+7]) {
				count++
			}
			if isValid(nums[i+2], nums[i+5], nums[i+8]) {
				count++
			}
		}
	}
	return count
}
