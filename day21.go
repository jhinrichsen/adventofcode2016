// all string operations check for their corresponding NOP action,
// and take all possible shortcuts.

package adventofcode2016

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func move(s string, p1, p2 int) string {
	if p1 == p2 {
		return s
	}
	original := s
	// remove at p1
	c := string(s[p1])
	s = s[:p1] + s[p1+1:]

	// insert
	s = s[:p2] + c + s[p2:]
	fmt.Printf("move(%s, %d, %d) = %s\n", original, p1, p2, s)
	return s
}

func swapPosition(s string, p1, p2 int) string {
	if p1 == p2 {
		return s
	}
	buf := []byte(s)
	buf[p1], buf[p2] = buf[p2], buf[p1]
	return string(buf)
}

func swapLetter(s string, a, b string) string {
	if a == b {
		return s
	}
	const (
		n        = 1
		noLetter = "_"
	)
	s = strings.Replace(s, a, noLetter, n)
	s = strings.Replace(s, b, a, n)
	s = strings.Replace(s, noLetter, b, n)
	return s
}

func reverse(s string, p1, p2 int) string {
	buf := []byte(s)
	slices.Reverse(buf[p1 : p2+1])
	return string(buf)
}

func rotateLeftN(s string, p int) string {
	l := len(s)
	p = p % l
	if p == 0 {
		return s
	}
	if p > l/2 {
		return rotateRightN(s, l-p)
	}
	for i := 0; i < p; i++ {
		s = s[1:] + s[0:1]
	}
	return s
}

func rotateRightN(s string, p int) string {
	l := len(s)
	p = p % l
	if p == 0 {
		return s
	}
	if p > l/2 {
		return rotateLeftN(s, l-p)
	}
	for i := 0; i < p; i++ {
		s = s[l-1:] + s[:l-1]
	}
	return s
}

func rotateLeftPos(s string, pos string) string {
	l := len(s)
	idx := strings.Index(s, pos)
	n := 1 + idx
	if idx >= 4 {
		n++
	}
	n = n % l
	/*
		if n > l/2 {
			return rotateRightN(s, l-n)
		}
	*/
	s2 := rotateLeftN(s, n)
	fmt.Printf("rotateLeftPos(%s, %s) = %s\n", s, pos, s2)
	return s2
}

func rotateRightPos(s string, pos string) string {
	l := len(s)
	idx := strings.Index(s, pos)
	n := 1 + idx
	if idx >= 4 {
		n++
	}
	n = n % l
	/*
		if n > l/2 {
			return rotateLeftN(s, l-n)
		}
	*/
	s2 := rotateRightN(s, n)
	fmt.Printf("rotateRightPos(%s, %s) = %s\n", s, pos, s2)
	return s2
}

func parseDay21(line string, scramble bool) func(s string) string {
	parts := strings.Fields(line)
	switch parts[0] {
	case "swap":
		a1 := parts[2]
		a2 := parts[5]
		if parts[1] == "letter" {
			return func(s string) string {
				if scramble {
					return swapLetter(s, a1, a2)
				}
				return swapLetter(s, a2, a1)
			}
		}
		p1, _ := strconv.Atoi(a1)
		p2, _ := strconv.Atoi(a2)
		return func(s string) string {
			if scramble {
				return swapPosition(s, p1, p2)
			}
			return swapPosition(s, p2, p1)
		}
	case "reverse":
		p1, _ := strconv.Atoi(parts[2])
		p2, _ := strconv.Atoi(parts[4])
		return func(s string) string {
			if scramble {
				return reverse(s, p1, p2)
			}
			return reverse(s, p2, p1)
		}
	case "rotate":
		var f func(string, int) string
		switch parts[1] {
		case "based":
			return func(s string) string {
				if scramble {
					return rotateRightPos(s, parts[6])
				}
				return rotateLeftPos(s, parts[6])
			}
		case "left":
			if scramble {
				f = rotateLeftN
			} else {
				f = rotateRightN
			}
		case "right":
			if scramble {
				f = rotateRightN
			} else {
				f = rotateLeftN
			}
		default:
			panic("bad index 1")
		}
		p1, _ := strconv.Atoi(parts[2])
		return func(s string) string {
			return f(s, p1)
		}
	case "move":
		p1, _ := strconv.Atoi(parts[2])
		p2, _ := strconv.Atoi(parts[5])
		return func(s string) string {
			if scramble {
				return move(s, p1, p2)
			}
			return move(s, p2, p1)
		}
	}
	return func(s string) string {
		return fmt.Sprintf("unknown command: %q", line)
	}
}

func Day21(password string, cmds []string, part1 bool) string {
	if !part1 {
		slices.Reverse(cmds)
	}
	for _, cmd := range cmds {
		password = parseDay21(cmd, part1)(password)
	}
	return password
}
