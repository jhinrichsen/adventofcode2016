package adventofcode2016

import "strings"

const (
	width  = 50
	height = 6
)

type day8 struct {
	grid          [][]bool
	width, height uint8
}

func newDay8(width, height uint8) day8 {
	var d day8
	d.width = width
	d.height = height
	d.grid = make([][]bool, d.width)
	for a := range d.grid {
		d.grid[a] = make([]bool, d.height)
	}
	return d
}

func (a *day8) rect(x, y uint) {
	for i := 0; i < int(x); i++ {
		for j := 0; j < int(y); j++ {
			a.grid[i][j] = true
		}
	}
}

func (a *day8) rotate(col bool, n uint, by uint) {
	var u0 uint8
	if col {
		tmp := make([]bool, height)
		for y := u0; y < a.height; y++ {
			tmp[y] = a.grid[n][y]
		}
		for src := u0; src < a.height; src++ {
			dst := (uint(src) + by) % uint(a.height)
			a.grid[n][dst] = tmp[src]
		}
	} else {
		tmp := make([]bool, a.width)
		for x := u0; x < a.width; x++ {
			tmp[x] = a.grid[x][n]
		}
		for src := u0; src < a.width; src++ {
			dst := (uint(src) + by) % uint(a.width)
			a.grid[dst][n] = tmp[src]
		}
	}
}

func (a day8) lit() (n uint) {
	for x := 0; x < int(a.width); x++ {
		for y := 0; y < int(a.height); y++ {
			if a.grid[x][y] {
				n++
			}
		}
	}
	return
}

func (a day8) String() string {
	var sb strings.Builder
	for y := 0; y < int(a.height); y++ {
		for x := 0; x < int(a.width); x++ {
			if a.grid[x][y] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')

			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// day8Callback will receive a 2D string representation of each and every step.
type day8Callback func(day8)

// Day8 returns the number of lit lights (lights in status 'on').
func Day8(screen day8, lines []string, part1 bool, f day8Callback) (uint, error) {
	for _, line := range lines {
		var col uint8

		rect := line[1] == 'e'
		var rx, ry uint

		rotate := line[1] == 'o'
		var rot byte // x or y
		var rotN uint
		var by uint

		var b byte
		add := func(n *uint) {
			if numeric(b) {
				*n *= 10
				*n += uint(b - '0')
			}
		}

		// state machine parser
		for j := 0; j < len(line); j++ {
			b = line[j]
			if b == ' ' {
				col++
				j++
				b = line[j]
			}
			if rect {
				// sample: rect 2x4
				if col == 1 {
					if b == 'x' {
						rx = ry
						ry = 0
					} else {
						add(&ry)
					}
				}
			} else if rotate {
				// sample: rotate column x=1 by 1
				if col == 2 {
					// once
					if rot == 0 {
						rot = b
					} else {
						add(&rotN)
					}
				} else if col == 4 {
					add(&by)
				}
			}
			// ignore other lines
		}

		if rect {
			screen.rect(rx, ry)
		} else if rotate {
			screen.rotate(rot == 'x', rotN, by)
		}

		if f != nil {
			f(screen)
		}
	}
	return screen.lit(), nil
}
