package adventofcode2016

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Opcode int

const (
	CPY Opcode = iota
	INC
	DEC
	JNZ
	TGL
)

type Operand struct {
	isRegister bool
	register   int // 0-3 for a-d
	immediate  int // literal value
}

type Instruction struct {
	opcode  Opcode
	arg1    Operand
	arg2    Operand
	hasArg2 bool
}

type Day23Puzzle struct {
	instructions []Instruction
}

func parseOperand(s string) Operand {
	if s >= "a" && s <= "d" {
		return Operand{isRegister: true, register: int(s[0] - 'a')}
	}
	val, _ := strconv.Atoi(s)
	return Operand{isRegister: false, immediate: val}
}

func parseOpcode(s string) Opcode {
	switch s {
	case "cpy":
		return CPY
	case "inc":
		return INC
	case "dec":
		return DEC
	case "jnz":
		return JNZ
	case "tgl":
		return TGL
	default:
		return CPY // fallback
	}
}

func NewDay23(lines []string) (Day23Puzzle, error) {
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		opcode := parseOpcode(fields[0])
		arg1 := parseOperand(fields[1])

		var arg2 Operand
		var hasArg2 bool
		if len(fields) >= 3 {
			arg2 = parseOperand(fields[2])
			hasArg2 = true
		}

		// Validate instruction format during parsing
		switch opcode {
		case CPY, JNZ:
			if !hasArg2 {
				return Day23Puzzle{}, fmt.Errorf("line %d: %s instruction missing second argument", i+1, fields[0])
			}
		case INC, DEC:
			if !arg1.isRegister {
				return Day23Puzzle{}, fmt.Errorf("line %d: %s instruction requires register argument, got %s", i+1, fields[0], fields[1])
			}
		case TGL:
			// TGL is valid with any arg1
		}

		instructions[i] = Instruction{
			opcode:  opcode,
			arg1:    arg1,
			arg2:    arg2,
			hasArg2: hasArg2,
		}
	}
	return Day23Puzzle{instructions: instructions}, nil
}

func Day23(puzzle Day23Puzzle, part1 bool) int {
	// create our working copy (toggle e.a. will change input)
	instructions := slices.Clone(puzzle.instructions)

	a := 12
	if part1 {
		a = 7
	}
	registers := [4]int{a, 0, 0, 0}

	getValue := func(op Operand) int {
		if op.isRegister {
			return registers[op.register]
		}
		return op.immediate
	}

	toggleInstruction := func(idx int) {
		if idx < 0 || idx >= len(instructions) {
			return // Outside program bounds
		}

		inst := &instructions[idx]
		switch inst.opcode {
		case INC:
			inst.opcode = DEC
		case DEC:
			inst.opcode = INC
		case TGL:
			inst.opcode = INC
		case JNZ:
			inst.opcode = CPY
		case CPY:
			inst.opcode = JNZ
		}
	}

	pc := 0
	for pc < len(instructions) {
		inst := instructions[pc]

		switch inst.opcode {
		case CPY:
			if inst.hasArg2 && inst.arg2.isRegister {
				value := getValue(inst.arg1)
				registers[inst.arg2.register] = value
			}
		case INC:
			registers[inst.arg1.register]++
		case DEC:
			registers[inst.arg1.register]--
		case JNZ:
			if inst.hasArg2 {
				value := getValue(inst.arg1)
				if value != 0 {
					offset := getValue(inst.arg2)
					pc += offset
					continue
				}
			}
		case TGL:
			offset := getValue(inst.arg1)
			target := pc + offset
			toggleInstruction(target)
		}
		pc++
	}

	return registers[0]
}
