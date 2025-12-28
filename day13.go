package adventofcode2016

import (
	"math/bits"
	"strings"
)

// isWall returns true if (x,y) is a wall given the favorite number
func isWall(x, y int, favorite uint) bool {
	if x < 0 || y < 0 {
		return true
	}
	ux, uy := uint(x), uint(y)
	val := ux*ux + 3*ux + 2*ux*uy + uy + uy*uy + favorite
	return bits.OnesCount(val)%2 == 1
}

type point13 struct {
	x, y int
}

// Day13Part1 returns minimum steps to go from src to dst using BFS.
func Day13Part1(favoriteNumber uint, src, dst complex64) (uint, error) {
	srcX, srcY := int(real(src)), int(imag(src))
	dstX, dstY := int(real(dst)), int(imag(dst))

	// BFS with visited map
	type state struct {
		x, y, dist int
	}

	visited := make(map[point13]bool)
	queue := []state{{srcX, srcY, 0}}
	visited[point13{srcX, srcY}] = true

	dirs := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.x == dstX && cur.y == dstY {
			return uint(cur.dist), nil
		}

		for _, d := range dirs {
			nx, ny := cur.x+d[0], cur.y+d[1]
			p := point13{nx, ny}
			if !visited[p] && !isWall(nx, ny, favoriteNumber) {
				visited[p] = true
				queue = append(queue, state{nx, ny, cur.dist + 1})
			}
		}
	}

	return 0, nil
}

// Day13Part2 returns number of locations reachable in at most `steps` steps.
func Day13Part2(favorite uint, src complex64, steps uint) uint {
	srcX, srcY := int(real(src)), int(imag(src))

	visited := make(map[point13]bool)
	type state struct {
		x, y, dist int
	}
	queue := []state{{srcX, srcY, 0}}
	visited[point13{srcX, srcY}] = true

	dirs := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.dist >= int(steps) {
			continue
		}

		for _, d := range dirs {
			nx, ny := cur.x+d[0], cur.y+d[1]
			p := point13{nx, ny}
			if !visited[p] && !isWall(nx, ny, favorite) {
				visited[p] = true
				queue = append(queue, state{nx, ny, cur.dist + 1})
			}
		}
	}

	return uint(len(visited))
}

// day13Grid for test compatibility
type day13Grid struct {
	favorite uint
}

func newDay13(favorite uint) *day13Grid {
	return &day13Grid{favorite: favorite}
}

func (g *day13Grid) String() string {
	const maxX, maxY = 10, 7
	var sb strings.Builder
	sb.WriteString("  0123456789\n")
	for y := 0; y < maxY; y++ {
		sb.WriteByte(byte(y) + '0')
		sb.WriteByte(' ')
		for x := 0; x < maxX; x++ {
			if isWall(x, y, g.favorite) {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		if y < maxY-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
