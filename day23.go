package adventofcode2016

import (
	"fmt"
	"strings"
)

// Day23 returns value of register a.
// In contrast to day 12, we do not have just compile time parameters, but also runtime parameters.
func Day23(lines []string, part1 bool) (int, error) {

	// type rtf func(a, b int) // run time function
	type rtf func(t bool)

	var registers [4]int
	var pc int // program counter
	words := make([]rtf, len(lines))

	isRegister := func(r byte) bool {
		return r >= 'a' && r <= 'd'
	}
	register := func(b byte) int {
		return int(b - 'a')
	}

	// increment register
	inc := func(r int) func(t bool) {
		return func(t bool) {
			registers[r]++
			pc++
		}
	}
	// decrement register
	dec := func(r int) func(t bool) {
		return func(t bool) {
			registers[r]--
			pc++
		}
	}
	tinc := func(ct bool, r int) func(rt bool) {
		ff := inc(r)
		ft := dec(r)

		var toggled bool
		return func(t bool) {
			if t {
				toggled = !toggled
				return
			}
			f := ff
			if toggled {
				f = ft
			}
			f(false)
		}
	}

	tdec := func(ct bool, r int) func(rt bool) {
		ff := dec(r)
		ft := inc(r)

		var toggled bool
		return func(t bool) {
			if t {
				toggled = !toggled
				return
			}
			f := ff
			if toggled {
				f = ft
			}
			f(false)
		}
	}

	tgl := func(r int) func() {
		return func() {
			n := registers[r]
			(words[pc+n])(true)
		}
	}

	ttgl := func(t bool, r int) func(t bool) {
		var toggled bool
		if t {
			toggled = !toggled
		}
		return func(t bool) {
			if toggled {
				inc(r)(false)
			} else {
				tgl(r)()
			}
		}
	}

	// copy immediate
	cpyi := func(n, r int) func(t bool) {
		return func(t bool) {
			registers[r] = n
			pc++
		}
	}
	// copy register
	cpyr := func(rx, ry int) func(t bool) {
		return func(t bool) {
			registers[ry] = registers[rx]
			pc++
		}
	}
	jnzi := func(x, n int) func(t bool) {
		return func(t bool) {
			// no jump
			if x == 0 {
				pc++
				return
			}
			// jump
			pc += n
		}
	}
	jnzr := func(r, n int) func(t bool) {
		return func(t bool) {
			// no jump
			if registers[r] == 0 {
				pc++
				return
			}
			// jump
			pc += n
		}
	}

	tcpy := func(t bool, a, b string) func(rt bool) {
		// figure out compile time part
		var ff, ft func(t bool)
		r0 := a[0]
		r1 := register(b[0])

		if isRegister(r0) {
			ft = jnzr(register(r0), r1)
		} else {
			ft = jnzi(toint(a), r1)
		}
		if isRegister(r0) {
			ff = cpyr(register(r0), r1)
		} else {
			ff = cpyi(toint(a), r1)
		}

		var toggled bool
		return func(t bool) {
			if t {
				toggled = !toggled
				return
			}
			f := ff
			if toggled {
				f = ft
			}
			f(false)
		}
	}

	if !part1 {
		registers[register('c')] = 1
	}

	// assemble phase

	var f func(t bool)
	for i, line := range lines {
		fs := strings.Fields(line)
		switch fs[0] {
		case "cpy":
			f = tcpy(false, fs[1], fs[2])
		case "inc":
			f = tinc(false, register(fs[1][0]))
		case "dec":
			f = tdec(false, register(fs[1][0]))
		case "jnz":
			r0 := fs[1][0]
			n := toint(fs[2])
			if isRegister(r0) {
				f = jnzr(register(r0), n)
			} else {
				f = jnzi(toint(fs[1]), n)
			}
		case "tgl":
			f = ttgl(false, register(fs[1][0]))

		default:
			return 0, fmt.Errorf("line %d: unknown instruction %q", pc, line)
		}
		words[i] = f
	}

	// run phase
	for pc < len(words) {
		(words[pc])(false)
	}
	return registers[register('a')], nil
}
