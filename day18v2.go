package adventofcode2016

import (
	"strings"
)

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
