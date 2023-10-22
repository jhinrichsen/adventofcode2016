package adventofcode2016

import (
	"strings"
)

const (
	safe = 46 // '.'
	trap = 94 // '^'
)

// derive uses ancestors to produce a new safe or trap.
// Its left and center tiles are traps, but its right tile is not.
// Its center and right tiles are traps, but its left tile is not.
// Only its left tile is a trap.
// Only its right tile is a trap.
func derive(left, right bool) bool {
	/* my working version
	b := (left && center && (!right)) ||
		(center && right && (!left)) ||
		(left && (!center) && (!right)) ||
		(right && (!left) && (!center))
	*/

	/* running the beast through a boolean algebra simplifier
	(a && b && (!c)) ||
	(b && c && (!a)) ||
	(a && (!b) && (!c)) ||
	(c && (!a) && (!b))

	=> (a && ~ c) || (~ a && c)
	*/

	// after a closer inspection, the expression
	// (left && (!right)) || ((!left) && right)
	// is a plain XOR

	// XOR the Go way:
	return left != right
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
		b := derive(left, right)
		var c byte = safe
		if b {
			c = trap
		}
		into[i] = c
	}
	return string(into)
}

func Day18(row string, count int) int {
	return Day18V1(row, count)
}

func Day18V1(row string, count int) int {
	var sum int
	// for i := count; i > 0; i-- {
	for i := 0; i < count; i++ {
		n := strings.Count(row, string(rune(safe)))
		sum += n
		row = next(row)
	}
	return sum
}

// safesAndTraps are bit representations of '.' and '^'.
// Bits are left and right padded with safe.
type safesAndTraps []bool

func (bits safesAndTraps) String() string {
	var sb strings.Builder
	for i := 1; i < len(bits)-1; i++ {
		var c byte
		if bits[i] {
			c = trap
		} else {
			c = safe
		}
		sb.WriteByte(c)
	}
	return sb.String()
}

func safes(a safesAndTraps) int {
	var n int
	for i := len(a) - 2; i > 0; i-- {
		if !a[i] {
			n++
		}
	}
	return n
}

func newSafesAndTraps(s string) safesAndTraps {
	bits := make([]bool, 1+len(s)+1) // left and right padding

	// we have two indices, loop over string index
	for si := len(s) - 1; si >= 0; si-- {
		bits[1+si] = s[si] == trap
	}
	return bits
}

func step(pred safesAndTraps, succ safesAndTraps) {
	// not sure if constants are lifted out of the loop
	// https://github.com/golang/go/issues/15808
	for i := len(pred) - 2; i > 0; i-- {
		succ[i] = derive(pred[i-1], pred[i+1])
	}
}

func Day18V2(row string, count int) int {
	a := newSafesAndTraps(row)
	b := make(safesAndTraps, len(a))

	var sum int
	for i := count; i > 0; i-- {
		sum += safes(a)
		step(a, b)
		a, b = b, a
	}

	return sum
}
