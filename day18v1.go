package adventofcode2016

import (
	"strings"
)

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
