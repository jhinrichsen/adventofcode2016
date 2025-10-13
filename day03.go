package adventofcode2016

import (
	"strconv"
	"strings"
)

// Angles is 3 for a Triangle.
const Angles = 3

// trianglePossible expects a triple and returns true if these form a triangle.
func trianglePossible(t1, t2, t3 uint) bool {
	if t3 < t2 {
		t2, t3 = t3, t2
	}
	if t2 < t1 {
		t1, t2 = t2, t1
	}
	if t3 < t2 {
		t2, t3 = t3, t2
	}
	return t1+t2 > t3
}

// Day03 returns number of triangles.
func Day03(lines []string, part1 bool) (possible uint) {
	chunk := func(lines []string) (g [Angles][Angles]uint) {
		for y := 0; y < Angles; y++ {
			fs := strings.Fields(lines[y])
			for x := 0; x < Angles; x++ {
				n, _ := strconv.Atoi(fs[x])
				g[y][x] = uint(n)
			}
		}
		return
	}
	rotate := func(from [Angles][Angles]uint) (into [Angles][Angles]uint) {
		for y := 0; y < len(from); y++ {
			for x := 0; x < len(from[0]); x++ {
				into[x][y] = from[y][x]
			}
		}
		return
	}
	for block := 0; block < len(lines); block += Angles {
		g := chunk(lines[block : block+Angles])
		if !part1 {
			g = rotate(g)
		}
		for y := 0; y < len(g); y++ {
			if trianglePossible(
				g[y][0],
				g[y][1],
				g[y][2],
			) {
				possible++
			}
		}
	}
	return
}
