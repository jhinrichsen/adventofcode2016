package adventofcode2016

import (
	"fmt"
	"strings"
)

// Day23 returns value of register a.
func Day23(lines []string, part1 bool) (int, error) {
	var registers [4]int

	isRegister := func(r byte) bool {
		return r >= 'a' && r <= 'd'
	}
	register := func(b byte) int {
		return int(b - 'a')
	}
	var pc int // program counter

	// copy immediate
	cpyi := func(n, r int) func() {
		return func() {
			registers[r] = n
			pc++
		}
	}
	// copy register
	cpyr := func(rx, ry int) func() {
		return func() {
			registers[ry] = registers[rx]
			pc++
		}
	}
	// increment register
	inc := func(r int) func() {
		return func() {
			registers[r]++
			pc++
		}
	}
	// decrement register
	dec := func(r int) func() {
		return func() {
			registers[r]--
			pc++
		}
	}
	jnzi := func(x, n int) func() {
		return func() {
			// no jump
			if x == 0 {
				pc++
				return
			}
			// jump
			pc += n
		}
	}
	jnzr := func(r, n int) func() {
		return func() {
			// no jump
			if registers[r] == 0 {
				pc++
				return
			}
			// jump
			pc += n
		}
	}
	if !part1 {
		registers[register('c')] = 1
	}

	// assemble phase

	var f func()
	words := make([]func(), len(lines))
	for i, line := range lines {
		fs := strings.Fields(line)
		switch fs[0] {
		case "cpy":
			r0 := fs[1][0]
			r1 := register(fs[2][0])
			if isRegister(r0) {
				f = cpyr(register(r0), r1)
			} else {
				f = cpyi(toint(fs[1]), r1)
			}
		case "inc":
			f = inc(register(fs[1][0]))
		case "dec":
			f = dec(register(fs[1][0]))
		case "jnz":
			r0 := fs[1][0]
			n := toint(fs[2])
			if isRegister(r0) {
				f = jnzr(register(r0), n)
			} else {
				f = jnzi(toint(fs[1]), n)
			}
		default:
			return 0, fmt.Errorf("line %d: unknown instruction %q", pc, line)
		}
		words[i] = f
	}

	// run phase
	for pc < len(words) {
		(words[pc])()
	}
	return registers[register('a')], nil
}
