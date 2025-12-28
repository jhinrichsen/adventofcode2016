package adventofcode2016

import (
	"fmt"
	"strings"
)

type stepfn21 func([]byte)

// Day21: Scrambled Letters and Hash
func Day21(cmds []string, phrase string, part1 bool) (string, error) {
	buf := []byte(phrase)
	tmp := make([]byte, len(buf)) // reusable temp buffer for rotations

	var scramblers []stepfn21
	var descramblers []stepfn21
	jt := newJumpTable(phrase)

	for _, cmd := range cmds {
		f1, f2, err := compile21(cmd, jt, tmp)
		if err != nil {
			return "", err
		}
		scramblers = append(scramblers, f1)
		descramblers = append(descramblers, f2)
	}

	if part1 {
		for i := 0; i < len(scramblers); i++ {
			scramblers[i](buf)
		}
	} else {
		for i := len(scramblers) - 1; i >= 0; i-- {
			descramblers[i](buf)
		}
	}
	return string(buf), nil
}

// compile21 reads one input line and creates a scrambler and descrambler.
// Uses inline parsing to avoid strings.Fields allocations.
func compile21(cmd string, jt jumpTable, tmp []byte) (stepfn21, stepfn21, error) {
	// Inline parsing helper
	parseNum := func(s string, start int) (int, int) {
		n := 0
		i := start
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
			i++
		}
		return n, i
	}

	skipWord := func(s string, start int) int {
		i := start
		// Skip leading spaces
		for i < len(s) && s[i] == ' ' {
			i++
		}
		// Skip the word
		for i < len(s) && s[i] != ' ' {
			i++
		}
		// Skip trailing spaces
		for i < len(s) && s[i] == ' ' {
			i++
		}
		return i
	}

	switch {
	case strings.HasPrefix(cmd, "swap position "):
		// "swap position X with position Y"
		i := 14 // after "swap position "
		p1, i := parseNum(cmd, i)
		i = skipWord(cmd, i) // skip "with"
		i = skipWord(cmd, i) // skip "position"
		p2, _ := parseNum(cmd, i)
		return func(b []byte) {
				b[p1], b[p2] = b[p2], b[p1]
			},
			func(b []byte) {
				b[p1], b[p2] = b[p2], b[p1]
			},
			nil

	case strings.HasPrefix(cmd, "swap letter "):
		// "swap letter X with letter Y"
		a := cmd[12]
		b := cmd[len(cmd)-1]
		return func(buf []byte) {
				swapLetterBytes(buf, a, b)
			},
			func(buf []byte) {
				swapLetterBytes(buf, a, b)
			},
			nil

	case strings.HasPrefix(cmd, "reverse positions "):
		// "reverse positions X through Y"
		i := 18 // after "reverse positions "
		p1, i := parseNum(cmd, i)
		i = skipWord(cmd, i) // skip "through"
		p2, _ := parseNum(cmd, i)
		return func(buf []byte) {
				reverseBytes(buf, p1, p2)
			},
			func(buf []byte) {
				reverseBytes(buf, p1, p2)
			},
			nil

	case strings.HasPrefix(cmd, "rotate based "):
		// "rotate based on position of letter X"
		c := cmd[len(cmd)-1]
		return func(buf []byte) {
				idx := indexByte(buf, c)
				n := jt[0][idx]
				rotateRightBytes(buf, n, tmp)
			},
			func(buf []byte) {
				idx := indexByte(buf, c)
				n := jt[1][idx]
				rotateLeftBytes(buf, n, tmp)
			},
			nil

	case strings.HasPrefix(cmd, "rotate left "):
		// "rotate left X step(s)"
		n, _ := parseNum(cmd, 12)
		return func(buf []byte) {
				rotateLeftBytes(buf, n, tmp)
			},
			func(buf []byte) {
				rotateRightBytes(buf, n, tmp)
			},
			nil

	case strings.HasPrefix(cmd, "rotate right "):
		// "rotate right X step(s)"
		n, _ := parseNum(cmd, 13)
		return func(buf []byte) {
				rotateRightBytes(buf, n, tmp)
			},
			func(buf []byte) {
				rotateLeftBytes(buf, n, tmp)
			},
			nil

	case strings.HasPrefix(cmd, "move position "):
		// "move position X to position Y"
		i := 14 // after "move position "
		p1, i := parseNum(cmd, i)
		i = skipWord(cmd, i) // skip "to"
		i = skipWord(cmd, i) // skip "position"
		p2, _ := parseNum(cmd, i)
		return func(buf []byte) {
				moveBytes(buf, p1, p2)
			},
			func(buf []byte) {
				moveBytes(buf, p2, p1)
			},
			nil
	}

	return nil, nil, fmt.Errorf("unknown command: %q", cmd)
}

func indexByte(buf []byte, c byte) int {
	for i, b := range buf {
		if b == c {
			return i
		}
	}
	return -1
}

func swapLetterBytes(buf []byte, a, b byte) {
	if a == b {
		return
	}
	ai, bi := -1, -1
	for i, c := range buf {
		if c == a {
			ai = i
		} else if c == b {
			bi = i
		}
		if ai >= 0 && bi >= 0 {
			break
		}
	}
	if ai >= 0 && bi >= 0 {
		buf[ai], buf[bi] = buf[bi], buf[ai]
	}
}

func reverseBytes(buf []byte, p1, p2 int) {
	for p1 < p2 {
		buf[p1], buf[p2] = buf[p2], buf[p1]
		p1++
		p2--
	}
}

func rotateLeftBytes(buf []byte, n int, tmp []byte) {
	l := len(buf)
	n = n % l
	if n == 0 {
		return
	}
	copy(tmp, buf)
	for i := 0; i < l; i++ {
		buf[i] = tmp[(i+n)%l]
	}
}

func rotateRightBytes(buf []byte, n int, tmp []byte) {
	l := len(buf)
	n = n % l
	if n == 0 {
		return
	}
	copy(tmp, buf)
	for i := 0; i < l; i++ {
		buf[i] = tmp[(i-n+l)%l]
	}
}

func moveBytes(buf []byte, from, to int) {
	if from == to {
		return
	}
	c := buf[from]
	if from < to {
		copy(buf[from:to], buf[from+1:to+1])
	} else {
		copy(buf[to+1:from+1], buf[to:from])
	}
	buf[to] = c
}

// Legacy functions for backwards compatibility with tests
func move(s string, p1, p2 int) string {
	buf := []byte(s)
	moveBytes(buf, p1, p2)
	return string(buf)
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
	buf := []byte(s)
	swapLetterBytes(buf, a[0], b[0])
	return string(buf)
}

func reverse(s string, p1, p2 int) string {
	buf := []byte(s)
	reverseBytes(buf, p1, p2)
	return string(buf)
}

func rotateLeftN(s string, p int) string {
	l := len(s)
	p = p % l
	if p == 0 {
		return s
	}
	return s[p:] + s[:p]
}

func rotateRightN(s string, p int) string {
	l := len(s)
	p = p % l
	if p == 0 {
		return s
	}
	return s[l-p:] + s[:l-p]
}

func rotateLeftPos(s string, pos string) string {
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
	return rotateRightN(s, n)
}

// compile is a compatibility wrapper for tests that use string-based stepfn
type stepfn func(string) string

func compile(cmd string, jt jumpTable) (stepfn, stepfn, error) {
	// Wrap byte-based functions to work with strings
	wrap := func(s string) string {
		buf := []byte(s)
		t := make([]byte, len(buf))
		f1, _, _ := compile21(cmd, jt, t)
		f1(buf)
		return string(buf)
	}
	wrapRev := func(s string) string {
		buf := []byte(s)
		t := make([]byte, len(buf))
		_, f2, _ := compile21(cmd, jt, t)
		f2(buf)
		return string(buf)
	}
	return wrap, wrapRev, nil
}
