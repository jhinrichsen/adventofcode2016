package adventofcode2016

func Day01(line string, part1 bool) uint {
	var x, y int
	var dx, dy int = 0, 1

	manhattan := func() uint {
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		return uint(x + y)
	}

	type pos struct{ x, y int }
	var positions map[pos]bool
	if !part1 {
		positions = make(map[pos]bool, len(line))
	}

	i := 0
	for i < len(line) {
		turn := line[i]
		i++

		var n int
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			n = n*10 + int(line[i]-'0')
			i++
		}

		if i < len(line) && line[i] == ',' {
			i += 2 // skip ", "
		}

		switch turn {
		case 'R':
			dx, dy = dy, -dx
		case 'L':
			dx, dy = -dy, dx
		}

		for j := 0; j < n; j++ {
			x += dx
			y += dy
			if !part1 {
				p := pos{x, y}
				if positions[p] {
					return manhattan()
				}
				positions[p] = true
			}
		}
	}
	return manhattan()
}
