package adventofcode2016

import (
	"strings"
)

const (
	safe = 46 // '.'
	trap = 94 // '^'
)

// Its left and center tiles are traps, but its right tile is not.
// Its center and right tiles are traps, but its left tile is not.
// Only its left tile is a trap.
// Only its right tile is a trap.
func isTrap(left, center, right bool) bool {
	/* my working version
	b := (left && center && (!right)) ||
		(center && right && (!left)) ||
		(left && (!center) && (!right)) ||
		(right && (!left) && (!center))
	*/

	/* running the beast through an boolean simplifier
	(a && b && (!c)) ||
	(b && c && (!a)) ||
	(a && (!b) && (!c)) ||
	(c && (!a) && (!b))

	=> (a && ~ c) || (~ a && c)
	*/

	return (left && (!right)) || ((!left) && right)
}

func next(row string) string {
	from := []byte(row)
	l := len(from)
	into := make([]byte, l)
	var left, right bool
	for i := 0; i < l; i++ {
		if i == 0 {
			left = false
		} else {
			left = from[i-1] == trap
		}
		if i+1 == l {
			right = false
		} else {
			right = from[i+1] == trap
		}
		center := from[i] == trap
		b := isTrap(left, center, right)
		var c byte = safe
		if b {
			c = trap
		}
		into[i] = c
	}
	return string(into)
}

func Day18(row string, count int) int {
	var n int
	for i := count; i > 0; i-- {
		n += strings.Count(row, string(rune(safe)))
		row = next(row)
	}
	return n
}
