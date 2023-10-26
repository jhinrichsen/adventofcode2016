package adventofcode2016

import (
	"fmt"
	"strings"
)

func newTraps(s string) []bool {
	bits := make([]bool, 1+len(s)+1) // left and right padding

	// we have two indices, loop over string index
	for si := len(s) - 1; si >= 0; si-- {
		bits[1+si] = s[si] == 94
	}
	return bits
}

func trapsAsString(bits []bool) string {
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

func Day18V3(row string, count int) int {
	a := newTraps(row)
	b := make([]bool, len(a))

	var sum int
	for ; count > 0; count-- {
		fmt.Printf("%s\n", trapsAsString(a))
		for i := len(a) - 2; i > 0; i-- {
			// count safes in current row
			if !a[i] {
				sum++
			}
			// determine next row
			b[i] = a[i-1] != a[i+1]
		}
		fmt.Printf("row %d: %d\n", count, sum)
		a, b = b, a
	}

	return sum
}
