package adventofcode2016

// Day02 solves day 2, both part 1 and part 2.
func Day02(input []byte, part1 bool) string {
	// Keypads as 5x5 grids (0 = invalid position)
	// Part 1: 3x3 keypad centered
	//     1 2 3
	//     4 5 6
	//     7 8 9
	var grid1 = [5][5]byte{
		{0, 0, 0, 0, 0},
		{0, '1', '2', '3', 0},
		{0, '4', '5', '6', 0},
		{0, '7', '8', '9', 0},
		{0, 0, 0, 0, 0},
	}
	// Part 2: diamond keypad
	//       1
	//     2 3 4
	//   5 6 7 8 9
	//     A B C
	//       D
	var grid2 = [5][5]byte{
		{0, 0, '1', 0, 0},
		{0, '2', '3', '4', 0},
		{'5', '6', '7', '8', '9'},
		{0, 'A', 'B', 'C', 0},
		{0, 0, 'D', 0, 0},
	}

	var grid *[5][5]byte
	var x, y int
	if part1 {
		grid = &grid1
		x, y = 2, 2 // start at '5'
	} else {
		grid = &grid2
		x, y = 0, 2 // start at '5'
	}

	// Count lines to pre-allocate
	lines := 0
	for _, b := range input {
		if b == '\n' {
			lines++
		}
	}
	if len(input) > 0 && input[len(input)-1] != '\n' {
		lines++
	}
	code := make([]byte, 0, lines)

	for _, b := range input {
		switch b {
		case '\n':
			code = append(code, grid[y][x])
		case 'U':
			if y > 0 && grid[y-1][x] != 0 {
				y--
			}
		case 'D':
			if y < 4 && grid[y+1][x] != 0 {
				y++
			}
		case 'L':
			if x > 0 && grid[y][x-1] != 0 {
				x--
			}
		case 'R':
			if x < 4 && grid[y][x+1] != 0 {
				x++
			}
		}
	}
	if len(input) > 0 && input[len(input)-1] != '\n' {
		code = append(code, grid[y][x])
	}

	if len(code) == 0 {
		return "0"
	}
	return string(code)
}
