package adventofcode2016

import (
	"fmt"
	"strconv"
	"strings"
)

// Day12 returns value of register a.
func Day12(lines []string, part1 bool) (int, error) {
	var registers [4]int

	isRegister := func(op string) bool {
		return len(op) == 1 && op[0] >= 'a' && op[0] <= 'd'
	}
	register := func(b byte) int {
		return int(b - 'a')
	}
	load := func(op string) (int, error) {
		if isRegister(op) {
			return registers[register(op[0])], nil
		}
		n, err := strconv.Atoi(op)
		if err != nil {
			return 0, fmt.Errorf("bad operand %q", op)
		}
		return n, nil
	}

	if !part1 {
		registers[register('c')] = 1
	}

	for pc := 0; pc < len(lines); {
		line := lines[pc]
		fs := strings.Fields(line)
		switch fs[0] {
		case "cpy":
			n, err := load(fs[1])
			if err != nil {
				return 0, fmt.Errorf("line %d: %w", pc, err)
			}
			registers[register(fs[2][0])] = n
			pc++
		case "inc":
			reg := register(fs[1][0])
			registers[reg]++
			pc++
		case "dec":
			reg := register(fs[1][0])
			registers[reg]--
			pc++
		case "jnz":
			n, err := load(fs[1])
			if err != nil {
				return 0, fmt.Errorf("line %d: %w", pc, err)
			}
			rel, err := strconv.Atoi(fs[2])
			if err != nil {
				return 0, fmt.Errorf("line %d: bad operand %q",
					pc, fs[2])
			}
			if n == 0 {
				pc++
			} else {
				// jump
				pc += rel
			}
		default:
			return 0, fmt.Errorf("line %d: unknown instruction %q", pc, line)
		}
	}
	return registers[0], nil
}
