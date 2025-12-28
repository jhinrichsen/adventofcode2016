package adventofcode2016

import "strings"

const (
	width  = 50
	height = 6
)

type day8 struct {
	grid          [width][height]bool // Fixed-size array, zero allocations
	w, h          int
	tmpCol        [height]bool // Reusable temp buffer for column rotation
	tmpRow        [width]bool  // Reusable temp buffer for row rotation
}

func newDay08(w, h uint8) day8 {
	return day8{w: int(w), h: int(h)}
}

func (a *day8) rect(x, y int) {
	for i := range x {
		for j := range y {
			a.grid[i][j] = true
		}
	}
}

func (a *day8) rotateCol(col, by int) {
	h := a.h
	// Copy column to temp buffer
	for y := range h {
		a.tmpCol[y] = a.grid[col][y]
	}
	// Rotate using temp buffer
	for src := range h {
		dst := (src + by) % h
		a.grid[col][dst] = a.tmpCol[src]
	}
}

func (a *day8) rotateRow(row, by int) {
	w := a.w
	// Copy row to temp buffer
	for x := range w {
		a.tmpRow[x] = a.grid[x][row]
	}
	// Rotate using temp buffer
	for src := range w {
		dst := (src + by) % w
		a.grid[dst][row] = a.tmpRow[src]
	}
}

func (a *day8) lit() uint {
	var n uint
	for x := range a.w {
		for y := range a.h {
			if a.grid[x][y] {
				n++
			}
		}
	}
	return n
}

func (a day8) String() string {
	var sb strings.Builder
	sb.Grow(a.h * (a.w + 1)) // Pre-allocate
	for y := range a.h {
		for x := range a.w {
			if a.grid[x][y] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// day8Callback will receive a 2D string representation of each and every step.
type day8Callback func(*day8)

// Day08 returns the number of lit lights (lights in status 'on').
func Day08(screen *day8, lines []string, part1 bool, f day8Callback) (uint, error) {
	for _, line := range lines {
		var col uint8

		isRect := line[1] == 'e'
		var rx, ry int

		isRotate := line[1] == 'o'
		var rot byte // x or y
		var rotN, by int

		// state machine parser
		for j := 0; j < len(line); j++ {
			b := line[j]
			if b == ' ' {
				col++
				j++
				b = line[j]
			}
			if isRect {
				// sample: rect 2x4
				if col == 1 {
					if b == 'x' {
						rx = ry
						ry = 0
					} else if b >= '0' && b <= '9' {
						ry = ry*10 + int(b-'0')
					}
				}
			} else if isRotate {
				// sample: rotate column x=1 by 1
				if col == 2 {
					// once
					if rot == 0 {
						rot = b
					} else if b >= '0' && b <= '9' {
						rotN = rotN*10 + int(b-'0')
					}
				} else if col == 4 {
					if b >= '0' && b <= '9' {
						by = by*10 + int(b-'0')
					}
				}
			}
		}

		if isRect {
			screen.rect(rx, ry)
		} else if isRotate {
			if rot == 'x' {
				screen.rotateCol(rotN, by)
			} else {
				screen.rotateRow(rotN, by)
			}
		}

		if f != nil {
			f(screen)
		}
	}
	return screen.lit(), nil
}
