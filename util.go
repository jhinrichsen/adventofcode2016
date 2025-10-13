package adventofcode2016

import "strconv"

// toint parses and returns a number as specifiec by strconv.Atoi or panic()s.
func toint(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
