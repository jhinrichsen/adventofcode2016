package adventofcode2016

type Day25Opcode int

const (
	Day25CPY Day25Opcode = iota
	Day25INC
	Day25DEC
	Day25JNZ
	Day25OUT
	Day25MUL // Peephole optimization: a += b * c
	Day25NOP
)

type Day25Operand struct {
	isRegister bool
	register   int
	immediate  int
}

type Day25Instruction struct {
	opcode  Day25Opcode
	arg1    Day25Operand
	arg2    Day25Operand
	arg3    Day25Operand // for MUL: target register
	hasArg2 bool
}

type Day25Puzzle struct {
	instructions []Day25Instruction
}

func parseDay25Operand(s string) Day25Operand {
	if len(s) == 1 && s[0] >= 'a' && s[0] <= 'd' {
		return Day25Operand{isRegister: true, register: int(s[0] - 'a')}
	}
	// Parse number inline
	neg := false
	i := 0
	if s[0] == '-' {
		neg = true
		i = 1
	}
	val := 0
	for ; i < len(s); i++ {
		val = val*10 + int(s[i]-'0')
	}
	if neg {
		val = -val
	}
	return Day25Operand{isRegister: false, immediate: val}
}

func parseDay25Opcode(s string) Day25Opcode {
	switch s {
	case "cpy":
		return Day25CPY
	case "inc":
		return Day25INC
	case "dec":
		return Day25DEC
	case "jnz":
		return Day25JNZ
	case "out":
		return Day25OUT
	}
	return Day25NOP
}

func NewDay25(lines []string) (Day25Puzzle, error) {
	instructions := make([]Day25Instruction, len(lines))

	// Parse instructions inline to avoid allocations
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		// Parse first field (opcode)
		j := 0
		for j < len(line) && line[j] != ' ' {
			j++
		}
		opcode := parseDay25Opcode(line[:j])

		// Skip space
		for j < len(line) && line[j] == ' ' {
			j++
		}

		// Parse first operand
		start := j
		for j < len(line) && line[j] != ' ' {
			j++
		}
		arg1 := parseDay25Operand(line[start:j])

		// Skip space
		for j < len(line) && line[j] == ' ' {
			j++
		}

		// Parse second operand if present
		var arg2 Day25Operand
		hasArg2 := false
		if j < len(line) {
			arg2 = parseDay25Operand(line[j:])
			hasArg2 = true
		}

		instructions[i] = Day25Instruction{
			opcode:  opcode,
			arg1:    arg1,
			arg2:    arg2,
			hasArg2: hasArg2,
		}
	}

	// Peephole optimization: detect multiplication patterns
	// Pattern 1: cpy X b; inc a; dec b; jnz b -2 → a += X
	// Pattern 2: cpy X c; cpy Y b; inc d; dec b; jnz b -2; dec c; jnz c -5 → d += X * Y
	for i := 0; i < len(instructions)-6; i++ {
		// Look for: cpy X c; cpy Y b; inc d; dec b; jnz b -2; dec c; jnz c -5
		if instructions[i].opcode == Day25CPY &&
			instructions[i+1].opcode == Day25CPY &&
			instructions[i+2].opcode == Day25INC &&
			instructions[i+3].opcode == Day25DEC &&
			instructions[i+4].opcode == Day25JNZ &&
			instructions[i+5].opcode == Day25DEC &&
			i+6 < len(instructions) && instructions[i+6].opcode == Day25JNZ {

			// Verify the pattern structure
			cReg := instructions[i].arg2   // target of first cpy (c)
			bReg := instructions[i+1].arg2 // target of second cpy (b)
			dReg := instructions[i+2].arg1 // inc target (d)

			if cReg.isRegister && bReg.isRegister && dReg.isRegister &&
				instructions[i+3].arg1.isRegister && instructions[i+3].arg1.register == bReg.register &&
				instructions[i+4].arg1.isRegister && instructions[i+4].arg1.register == bReg.register &&
				instructions[i+5].arg1.isRegister && instructions[i+5].arg1.register == cReg.register &&
				instructions[i+6].arg1.isRegister && instructions[i+6].arg1.register == cReg.register {

				// Check jump offsets
				if !instructions[i+4].arg2.isRegister && instructions[i+4].arg2.immediate == -2 &&
					!instructions[i+6].arg2.isRegister && instructions[i+6].arg2.immediate == -5 {

					// Replace with MUL: d += X * Y, clear b and c
					instructions[i] = Day25Instruction{
						opcode: Day25MUL,
						arg1:   instructions[i].arg1,   // X (multiplier from c)
						arg2:   instructions[i+1].arg1, // Y (multiplier from b)
						arg3:   dReg,                   // target register
					}
					// NOP out the rest
					for k := i + 1; k <= i+6; k++ {
						instructions[k] = Day25Instruction{opcode: Day25NOP}
					}
				}
			}
		}
	}

	return Day25Puzzle{instructions: instructions}, nil
}

func Day25(puzzle Day25Puzzle) uint {
	for a := uint(1); a < 1000; a++ {
		if generatesClock(puzzle, a) {
			return a
		}
	}
	return 0
}

func generatesClock(puzzle Day25Puzzle, initialA uint) bool {
	registers := [4]int{int(initialA), 0, 0, 0}
	outputCount := 0
	expectedNext := 0
	maxOutputs := 20
	maxInstructions := 100000

	pc := 0
	instructionCount := 0

	for pc < len(puzzle.instructions) && outputCount < maxOutputs && instructionCount < maxInstructions {
		inst := puzzle.instructions[pc]
		instructionCount++

		switch inst.opcode {
		case Day25CPY:
			if inst.hasArg2 && inst.arg2.isRegister {
				var value int
				if inst.arg1.isRegister {
					value = registers[inst.arg1.register]
				} else {
					value = inst.arg1.immediate
				}
				registers[inst.arg2.register] = value
			}
		case Day25INC:
			registers[inst.arg1.register]++
		case Day25DEC:
			registers[inst.arg1.register]--
		case Day25JNZ:
			if inst.hasArg2 {
				var value int
				if inst.arg1.isRegister {
					value = registers[inst.arg1.register]
				} else {
					value = inst.arg1.immediate
				}
				if value != 0 {
					var offset int
					if inst.arg2.isRegister {
						offset = registers[inst.arg2.register]
					} else {
						offset = inst.arg2.immediate
					}
					pc += offset
					continue
				}
			}
		case Day25OUT:
			var value int
			if inst.arg1.isRegister {
				value = registers[inst.arg1.register]
			} else {
				value = inst.arg1.immediate
			}
			if value != expectedNext {
				return false
			}
			expectedNext = 1 - expectedNext
			outputCount++
		case Day25MUL:
			// a += X * Y, where X and Y are arg1 and arg2
			var x, y int
			if inst.arg1.isRegister {
				x = registers[inst.arg1.register]
			} else {
				x = inst.arg1.immediate
			}
			if inst.arg2.isRegister {
				y = registers[inst.arg2.register]
			} else {
				y = inst.arg2.immediate
			}
			registers[inst.arg3.register] += x * y
		case Day25NOP:
			// Do nothing
		}
		pc++
	}

	return outputCount >= 10
}
