package adventofcode2016

import (
	"fmt"
	"strconv"
	"strings"
)

type Day25Opcode int

const (
	Day25CPY Day25Opcode = iota
	Day25INC
	Day25DEC
	Day25JNZ
	Day25OUT
)

type Day25Operand struct {
	isRegister bool
	register   int // 0-3 for a-d
	immediate  int // literal value
}

type Day25Instruction struct {
	opcode  Day25Opcode
	arg1    Day25Operand
	arg2    Day25Operand
	hasArg2 bool
}

type Day25Puzzle struct {
	instructions []Day25Instruction
}

func parseDay25Operand(s string) (Day25Operand, error) {
	if s >= "a" && s <= "d" {
		return Day25Operand{isRegister: true, register: int(s[0] - 'a')}, nil
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		return Day25Operand{}, fmt.Errorf("invalid operand %q: %w", s, err)
	}
	return Day25Operand{isRegister: false, immediate: val}, nil
}

func parseDay25Opcode(s string) (Day25Opcode, error) {
	switch s {
	case "cpy":
		return Day25CPY, nil
	case "inc":
		return Day25INC, nil
	case "dec":
		return Day25DEC, nil
	case "jnz":
		return Day25JNZ, nil
	case "out":
		return Day25OUT, nil
	default:
		return 0, fmt.Errorf("unknown opcode: %q", s)
	}
}

func NewDay25(lines []string) (Day25Puzzle, error) {
	instructions := make([]Day25Instruction, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		opcode, err := parseDay25Opcode(fields[0])
		if err != nil {
			return Day25Puzzle{}, fmt.Errorf("line %d: %w", i+1, err)
		}
		arg1, err := parseDay25Operand(fields[1])
		if err != nil {
			return Day25Puzzle{}, fmt.Errorf("line %d: %w", i+1, err)
		}

		var arg2 Day25Operand
		var hasArg2 bool
		if len(fields) >= 3 {
			arg2, err = parseDay25Operand(fields[2])
			if err != nil {
				return Day25Puzzle{}, fmt.Errorf("line %d: %w", i+1, err)
			}
			hasArg2 = true
		}

		// Validate instruction format during parsing
		switch opcode {
		case Day25CPY, Day25JNZ:
			if !hasArg2 {
				return Day25Puzzle{}, fmt.Errorf("line %d: %s instruction missing second argument", i+1, fields[0])
			}
		case Day25INC, Day25DEC:
			if !arg1.isRegister {
				return Day25Puzzle{}, fmt.Errorf("line %d: %s instruction requires register argument, got %s", i+1, fields[0], fields[1])
			}
		case Day25OUT:
			// OUT is valid with any arg1
		}

		instructions[i] = Day25Instruction{
			opcode:  opcode,
			arg1:    arg1,
			arg2:    arg2,
			hasArg2: hasArg2,
		}
	}
	return Day25Puzzle{instructions: instructions}, nil
}

func Day25(puzzle Day25Puzzle) uint {
	// Try different values of register 'a' starting from 1
	for a := uint(1); a < 1000; a++ {
		if generatesClock(puzzle, a) {
			return a
		}
	}
	return 0 // Should not reach here if input is valid
}

func generatesClock(puzzle Day25Puzzle, initialA uint) bool {
	registers := [4]int{int(initialA), 0, 0, 0}
	output := make([]int, 0, 20) // Collect first 20 outputs
	maxOutputs := 20
	maxInstructions := 100000 // Prevent infinite loops

	getValue := func(op Day25Operand) int {
		if op.isRegister {
			return registers[op.register]
		}
		return op.immediate
	}

	pc := 0
	instructionCount := 0

	for pc < len(puzzle.instructions) && len(output) < maxOutputs && instructionCount < maxInstructions {
		inst := puzzle.instructions[pc]
		instructionCount++

		switch inst.opcode {
		case Day25CPY:
			if inst.hasArg2 && inst.arg2.isRegister {
				value := getValue(inst.arg1)
				registers[inst.arg2.register] = value
			}
		case Day25INC:
			registers[inst.arg1.register]++
		case Day25DEC:
			registers[inst.arg1.register]--
		case Day25JNZ:
			if inst.hasArg2 {
				value := getValue(inst.arg1)
				if value != 0 {
					offset := getValue(inst.arg2)
					pc += offset
					continue
				}
			}
		case Day25OUT:
			value := getValue(inst.arg1)
			output = append(output, value)
		}
		pc++
	}

	// Check if output alternates between 0 and 1
	if len(output) < 10 {
		return false // Need at least some output to verify pattern
	}

	for i, val := range output {
		expected := i % 2
		if val != expected {
			return false
		}
	}

	return true
}