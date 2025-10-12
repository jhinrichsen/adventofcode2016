// all string operations check for their corresponding NOP action,
// and take all possible shortcuts.

package adventofcode2016

import (
	"fmt"
	"slices"
	"strings"
)

type stepfn func(string) string

// Day21: Scrambled Letters and Hash
func Day21(cmds []string, phrase string, part1 bool) (string, error) {
	var scramblers []stepfn
	var descramblers []stepfn // list of reverse scramble stepper up to the first ambiguous one
	jt := newJumpTable(phrase)
	for _, cmd := range cmds {
		f1, f2, err := compile(cmd, jt)
		if err != nil {
			return "", err
		}
		scramblers = append(scramblers, f1)
		descramblers = append(descramblers, f2)
	}

	if part1 {
		for i := 0; i < len(scramblers); i++ {
			phrase = scramblers[i](phrase)
		}
	} else {
		for i := len(scramblers) - 1; i >= 0; i-- {
			phrase = descramblers[i](phrase)
		}
	}
	return phrase, nil
}

// compile reads one input line, and creates a scrambler, and an optional descrambler.
func compile(cmd string, jt jumpTable) (stepfn, stepfn, error) {
	tokens := strings.Fields(cmd)
	num := func(n int) int {
		return toint(tokens[n])
	}
	switch tokens[0] {
	case "swap":
		a1 := tokens[2]
		a2 := tokens[5]
		if tokens[1] == "letter" {
			return func(s string) string {
					return swapLetter(s, a1, a2)
				},
				func(s string) string {
					return swapLetter(s, a2, a1)
				},
				nil
		}
		p1 := toint(a1)
		p2 := toint(a2)
		return func(s string) string {
				return swapPosition(s, p1, p2)
			},
			func(s string) string {
				return swapPosition(s, p2, p1)
			},
			nil
	case "reverse":
		p1 := num(2)
		p2 := num(4)
		// reverse is reverse no matter from what point of view
		return func(s string) string {
				return reverse(s, p1, p2)
			},
			func(s string) string {
				return reverse(s, p1, p2)
			},
			nil
	case "rotate":
		switch tokens[1] {
		case "based":
			c := tokens[6]
			return func(s string) string {
					idx := strings.Index(s, c)
					n := jt[0][idx]
					return rotateRightN(s, n)
				},
				func(s string) string {
					idx := strings.Index(s, c)
					n := jt[1][idx]
					return rotateLeftN(s, n)
				},
				nil
		case "left":
			return func(s string) string {
					return rotateLeftN(s, num(2))
				},
				func(s string) string {
					return rotateRightN(s, num(2))
				},
				nil
		case "right":
			return func(s string) string {
					return rotateRightN(s, num(2))
				},
				func(s string) string {
					return rotateLeftN(s, num(2))
				},
				nil
		}
		panic("bad index 1")
	case "move":
		p1 := num(2)
		p2 := num(5)
		return func(s string) string {
				return move(s, p1, p2)
			},
			func(s string) string {
				return move(s, p2, p1)

			},
			nil
	}
	return nil, nil, fmt.Errorf("unknown command: %q", cmd)
}

func move(s string, p1, p2 int) string {
	if p1 == p2 {
		return s
	}
	// remove at p1
	c := string(s[p1])
	s = s[:p1] + s[p1+1:]

	// insert
	s = s[:p2] + c + s[p2:]
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
	/*
		if p > l/2 {
			return rotateRightN(s, l-p)
		}
	*/
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
	/*
		if p > l/2 {
			return rotateLeftN(s, l-p)
		}
	*/
	for i := 0; i < p; i++ {
		s = s[l-1:] + s[:l-1]
	}
	return s
}

func rotateLeftPos(s string, pos string) string {
	// This implements "rotate based on position" - same as rotateRightPos
	// The "Left" in the name is misleading; both implement the same AoC operation
	l := len(s)
	idx := strings.Index(s, pos)
	n := 1 + idx
	if idx >= 4 {
		n++
	}
	n = n % l
	return rotateRightN(s, n)
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
	return s2
}
