package adventofcode2016

import (
	"fmt"
	"math/bits"
	"strings"

	"github.com/beefsack/go-astar"
)

type cubicle struct {
	x, y int
	wall bool   // wall or open space
	area *day13 // back reference to complete area

	// included in part 2
	floodfilled bool
}

func (a *cubicle) PathNeighbors() []astar.Pather {
	var hood []astar.Pather
	// north
	y := a.y - 1
	if y >= 0 && !a.area[a.x][y].wall {
		hood = append(hood, &a.area[a.x][y])
	}
	// east
	x := a.x + 1
	if x < len(a.area) && !a.area[x][a.y].wall {
		hood = append(hood, &a.area[x][a.y])
	}
	// south
	y = a.y + 1
	if y < len(a.area[0]) && !a.area[a.x][y].wall {
		hood = append(hood, &a.area[a.x][y])
	}
	// west
	x = a.x - 1
	if x >= 0 && !a.area[x][a.y].wall {
		hood = append(hood, &a.area[x][a.y])
	}
	return hood
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (a *cubicle) PathNeighborCost(to astar.Pather) float64 {
	return 1.0
}

func (a *cubicle) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*cubicle)
	absX := toT.x - a.x
	if absX < 0 {
		absX = -absX
	}
	absY := toT.y - a.y
	if absY < 0 {
		absY = -absY
	}
	d := float64(absX + absY)
	return d
}

type day13 [51][51]cubicle // part2 requires depth 50

// filled returns number of filled cubicles.
func (a day13) filled() uint {
	sum := uint(0)
	for y := 0; y < len(a); y++ {
		for x := 0; x < len(a[0]); x++ {
			if a[x][y].floodfilled {
				sum++
			}
		}
	}
	return sum
}

// String return 10x6 string representation.
func (a day13) String() string {
	const (
		maxX = 10
		maxY = 7

		openSpace = '.'
		wall      = '#'
		newline   = '\n'
		separator = ' '
	)
	var sb strings.Builder
	sb.WriteString("  0123456789\n")
	for y := 0; y < maxY; y++ {
		sb.WriteByte(byte(y) + '0')
		sb.WriteByte(separator)
		for x := 0; x < maxX; x++ {
			if a[x][y].wall {
				sb.WriteByte(wall)
			} else {
				sb.WriteByte(openSpace)
			}
		}
		if y < maxY-1 {
			sb.WriteByte(newline)
		}
	}
	return sb.String()
}

// Day13Part1 returns minimum steps to go from (1,1) to (31,39).
func Day13Part1(favoriteNumber uint, src, dst complex64) (uint, error) {
	d := newDay13(favoriteNumber)
	from := &d[int(real(src))][int(imag(src))]
	to := &d[int(real(dst))][int(imag(dst))]
	_, distance, found := astar.Path(from, to)

	if !found {
		return 0, fmt.Errorf("no possible path")
	}
	return uint(distance), nil
}

// newDay13 returns a pointer, so that the result will be the same pointer
// as any embedded cubicle's area. When using return by value, taking the
// address of day13 for starting and end point will break A*, which operates
// on internal area pointer.
func newDay13(favoriteNumber uint) *day13 {
	f := func(x, y uint) uint {
		return x*x + 3*x + 2*x*y + y + y*y
	}
	bitsSet := func(x uint) uint8 {
		return uint8(bits.OnesCount(x))
	}
	odd := func(x uint8) bool {
		return x%2 != 0
	}
	var d day13
	for x := 0; x < len(d); x++ {
		for y := 0; y < len(d[0]); y++ {
			d[x][y].x = x
			d[x][y].y = y
			ux, uy := uint(x), uint(y)
			d[x][y].wall = odd(bitsSet(favoriteNumber + f(ux, uy)))
			d[x][y].area = &d
		}
	}
	return &d
}

// Day13Part2 returns number of fields flooded for depth = 50 beginning from
// (1,1).
func Day13Part2(favorite uint, src complex64, steps uint) uint {
	d := newDay13(favorite)
	scratchpad := make(map[cubicle]bool)
	scratchpad[d[1][1]] = true
	for step := uint(0); step <= steps; step++ {
		neighbours := make(map[cubicle]bool)
		for k := range scratchpad {
			d[k.x][k.y].floodfilled = true
			hood := k.PathNeighbors()
			for _, h := range hood {
				c := h.(*cubicle)
				if !c.floodfilled {
					neighbours[*c] = true
				}
			}
			delete(scratchpad, k)
		}
		scratchpad = neighbours
	}
	return d.filled()
}
