package adventofcode2016

import (
	"fmt"
	"strconv"
	"strings"
)

// bot is also used for outputs.
type bot struct {
	ID     string
	n      int // input counter, type int because of index into []input
	input  [2]uint
	output [2]string // ID of other bots/ outputs, low is low index (0)
}

func (a *bot) send(value uint) {
	a.input[a.n] = value
	a.n++
}

func (a bot) active() bool {
	return a.n == 2
}

// Day10 solves day 10. For part 2, value1 and value2 are ignored.
func Day10(lines []string, part1 bool, value1, value2 uint) (uint, error) {
	// invariant: value1 < value2
	if value1 > value2 {
		value1, value2 = value2, value1
	}
	bots := make(map[string]bot)
	active := make(map[string]bool)
	for i, line := range lines {
		fs := strings.Fields(line)
		if fs[0] == "value" {
			// parse "value 5 goes to bot 2"
			vi, err := strconv.Atoi(fs[1])
			if err != nil {
				return 0, fmt.Errorf("line %d: not a number: %q",
					i, fs[1])
			}
			v := uint(vi)
			ID := fs[4] + fs[5]
			if len(ID) == 0 {
				return 0, fmt.Errorf("line %d: missing ID in %q",
					i, line)
			}
			b := bots[ID]
			b.ID = ID
			b.send(v)
			bots[ID] = b
			if b.active() {
				active[ID] = true
			}
		} else {
			// parse "bot 2 gives low to bot 1 and high to bot 0"
			if len(fs) != 12 {
				msg := "want 12 columns but got %d: %q"
				return 0,
					fmt.Errorf(msg, len(fs), line)
			}
			ID := fs[0] + fs[1]
			if len(ID) == 0 {
				return 0, fmt.Errorf("line %d: missing ID in %q",
					i, line)
			}
			b := bots[ID]
			b.ID = ID
			ID0 := fs[5] + fs[6]
			ID1 := fs[10] + fs[11]
			b.output = [...]string{
				ID0,
				ID1,
			}
			bots[ID] = b
			// register output bots
			if _, ok := bots[ID0]; !ok {
				bots[ID0] = bot{ID: ID0}
			}
			if _, ok := bots[ID1]; !ok {
				bots[ID1] = bot{ID: ID1}
			}
		}
	}

	for len(active) > 0 {
		var k string
		for k = range active {
			break
		}
		delete(active, k)
		b := bots[k]
		l, h := b.input[0], b.input[1]
		if l > h {
			l, h = h, l
		}

		// this bot?
		if part1 && l == value1 && h == value2 {
			x, err := strconv.Atoi(b.ID[3:])
			if err != nil {
				return 0, fmt.Errorf("error converting bot ID "+
					"%q to number", b.ID)
			}
			return uint(x), nil
		}
		if low, ok := bots[b.output[0]]; ok {
			low.send(l)
			bots[low.ID] = low
			if low.active() {
				active[low.ID] = true
			}
		} else {
			return 0, fmt.Errorf("%q has unknown low bot %q",
				b.ID, b.output[0])
		}
		if high, ok := bots[b.output[1]]; ok {
			high.send(h)
			bots[high.ID] = high
			if high.active() {
				active[high.ID] = true
			}
		} else {
			return 0, fmt.Errorf("%q has unknown high bot %q",
				b.ID, b.output[1])
		}
	}
	product := uint(1)
	for _, ID := range []string{"output0", "output1", "output2"} {
		b := bots[ID]
		for i := 0; i < b.n; i++ {
			product *= b.input[i]
		}
	}
	return product, nil
}
