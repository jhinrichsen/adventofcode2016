package adventofcode2016

import (
	"math"
)

// Day22: Grid Computing
func Day22(lines []string, part1 bool) (uint, error) {
	// Inline parsing: 'node-x1-y2' -> 1, 2
	parseNode := func(s string) (int, int) {
		// Format: /dev/grid/node-xN-yM
		// Find "node-x" then parse numbers
		i := 0
		for i < len(s) && s[i] != 'x' {
			i++
		}
		i++ // skip 'x'
		x := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			x = x*10 + int(s[i]-'0')
			i++
		}
		i++ // skip '-'
		i++ // skip 'y'
		y := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			y = y*10 + int(s[i]-'0')
			i++
		}
		return x, y
	}

	// Parse human-readable size like "72T" or "20T"
	// Returns value in TB (we can use TB as base unit since all values use T)
	parseSize := func(s string) uint64 {
		n := uint64(0)
		for i := 0; i < len(s)-1; i++ { // -1 to skip unit suffix
			n = n*10 + uint64(s[i]-'0')
		}
		return n
	}

	// Skip whitespace, return position of next non-space
	skipSpace := func(s string, i int) int {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		return i
	}

	// Find end of field (next space or end of string)
	fieldEnd := func(s string, i int) int {
		for i < len(s) && s[i] != ' ' {
			i++
		}
		return i
	}

	type df struct {
		used, avail uint64
	}
	estimatedNodes := int(math.Sqrt(float64(len(lines))))
	nodes := make([]df, 0, estimatedNodes)

	var dimX, dimY int
	var emptyX, emptyY int

	for _, line := range lines {
		if len(line) == 0 || line[0] != '/' {
			continue
		}

		// Parse: /dev/grid/node-xN-yM   Size  Used  Avail  Use%
		i := 0
		end := fieldEnd(line, i)
		x, y := parseNode(line[i:end])

		dimX = max(dimX, x)
		dimY = max(dimY, y)

		// Skip to Used field (skip Size)
		i = skipSpace(line, end)
		i = fieldEnd(line, i) // skip Size
		i = skipSpace(line, i)
		end = fieldEnd(line, i)
		used := parseSize(line[i:end])

		// Parse Avail
		i = skipSpace(line, end)
		end = fieldEnd(line, i)
		avail := parseSize(line[i:end])

		if used == 0 {
			emptyX, emptyY = x, y
		}

		nodes = append(nodes, df{used, avail})
	}

	// dimension is one larger than max index
	dimX++
	dimY++
	_ = emptyX
	_ = emptyY

	// part 1: count viable pairs
	var viable uint
	for i, na := range nodes {
		if na.used == 0 {
			continue
		}
		for j, nb := range nodes {
			if i == j {
				continue
			}
			if na.used <= nb.avail {
				viable++
			}
		}
	}
	if part1 {
		return viable, nil
	}

	// Part 2: Count moves to get data from top-right to top-left
	type path struct {
		dx, dy int
		n      int
	}
	var (
		paths = []path{
			{-1, 0, 4},  // left 4
			{0, -1, 22}, // up 22
			{1, 0, 22},  // right 22
		}
	)

	// repeat sequence to move red hole to the left
	for i := 0; i < dimX-2; i++ {
		paths = append(paths, []path{
			{0, 1, 1},   // down 1
			{-1, 0, 2},  // left 2
			{0, -1, 1},  // up 1
			{1, 0, 1},   // right 1
		}...)
	}

	var count uint
	for i := 0; i < len(paths); i++ {
		count += uint(paths[i].n)
	}
	return count, nil
}
