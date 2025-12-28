package adventofcode2016

// bot10 is used for bots and outputs in Day10.
type bot10 struct {
	n         int      // input counter
	input     [2]uint  // chip values
	lowID     int      // destination for low chip (negative = output)
	highID    int      // destination for high chip (negative = output)
	hasOutput bool     // whether output destinations are set
}

// Day10 solves day 10. For part 2, value1 and value2 are ignored.
func Day10(lines []string, part1 bool, value1, value2 uint) (uint, error) {
	// invariant: value1 < value2
	if value1 > value2 {
		value1, value2 = value2, value1
	}

	bots := make(map[int]*bot10)
	outputs := make(map[int]*bot10)
	var active []int // bot IDs that have 2 chips

	getBot := func(id int) *bot10 {
		if b, ok := bots[id]; ok {
			return b
		}
		b := &bot10{}
		bots[id] = b
		return b
	}

	getOutput := func(id int) *bot10 {
		if b, ok := outputs[id]; ok {
			return b
		}
		b := &bot10{}
		outputs[id] = b
		return b
	}

	for _, line := range lines {
		if line[0] == 'v' {
			// parse "value 5 goes to bot 2"
			i := 6 // skip "value "
			var val uint
			for line[i] != ' ' {
				val = val*10 + uint(line[i]-'0')
				i++
			}
			// skip " goes to bot "
			i += 13
			var botID int
			for i < len(line) {
				botID = botID*10 + int(line[i]-'0')
				i++
			}
			b := getBot(botID)
			b.input[b.n] = val
			b.n++
			if b.n == 2 && b.hasOutput {
				active = append(active, botID)
			}
		} else {
			// parse "bot 2 gives low to bot 1 and high to bot 0"
			// or   "bot 2 gives low to output 1 and high to bot 0"
			i := 4 // skip "bot "
			var botID int
			for line[i] != ' ' {
				botID = botID*10 + int(line[i]-'0')
				i++
			}
			// skip " gives low to "
			i += 14
			lowIsOutput := line[i] == 'o'
			if lowIsOutput {
				i += 7 // skip "output "
			} else {
				i += 4 // skip "bot "
			}
			var lowID int
			for line[i] != ' ' {
				lowID = lowID*10 + int(line[i]-'0')
				i++
			}
			// skip " and high to "
			i += 13
			highIsOutput := line[i] == 'o'
			if highIsOutput {
				i += 7 // skip "output "
			} else {
				i += 4 // skip "bot "
			}
			var highID int
			for i < len(line) {
				highID = highID*10 + int(line[i]-'0')
				i++
			}

			b := getBot(botID)
			if lowIsOutput {
				b.lowID = -(lowID + 1) // negative = output, +1 to distinguish from 0
			} else {
				b.lowID = lowID + 1 // positive = bot, +1 to distinguish from unset
			}
			if highIsOutput {
				b.highID = -(highID + 1)
			} else {
				b.highID = highID + 1
			}
			b.hasOutput = true
			if b.n == 2 {
				active = append(active, botID)
			}
		}
	}

	for len(active) > 0 {
		botID := active[len(active)-1]
		active = active[:len(active)-1]

		b := bots[botID]
		l, h := b.input[0], b.input[1]
		if l > h {
			l, h = h, l
		}

		// this bot?
		if part1 && l == value1 && h == value2 {
			return uint(botID), nil
		}

		// send low
		if b.lowID < 0 {
			// output
			out := getOutput(-(b.lowID + 1))
			out.input[out.n] = l
			out.n++
		} else {
			// bot
			dest := getBot(b.lowID - 1)
			dest.input[dest.n] = l
			dest.n++
			if dest.n == 2 && dest.hasOutput {
				active = append(active, b.lowID-1)
			}
		}

		// send high
		if b.highID < 0 {
			// output
			out := getOutput(-(b.highID + 1))
			out.input[out.n] = h
			out.n++
		} else {
			// bot
			dest := getBot(b.highID - 1)
			dest.input[dest.n] = h
			dest.n++
			if dest.n == 2 && dest.hasOutput {
				active = append(active, b.highID-1)
			}
		}
	}

	product := uint(1)
	for i := 0; i < 3; i++ {
		if out, ok := outputs[i]; ok {
			for j := 0; j < out.n; j++ {
				product *= out.input[j]
			}
		}
	}
	return product, nil
}
