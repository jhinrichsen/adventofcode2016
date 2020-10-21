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

// Day10 solves day 10 part 1.
func Day10(lines []string, value1, value2 uint) (uint, error) {
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
			b := bots[ID]
			b.ID = ID
			b.output = [...]string{
				fs[5] + fs[6],
				fs[10] + fs[11],
			}
			bots[ID] = b
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
		if l == value1 && h == value2 {
			x, err := strconv.Atoi(b.ID[3:])
			if err != nil {
				return 0, fmt.Errorf("error converting bot ID "+
					"%q to number", b.ID)
			}
			return uint(x), nil
		}
		low := bots[b.output[0]]
		low.send(l)
		bots[low.ID] = low
		if low.active() {
			active[low.ID] = true
		}
		high := bots[b.output[1]]
		high.send(h)
		bots[high.ID] = high
		if high.active() {
			active[high.ID] = true
		}
	}
	return 0, fmt.Errorf("no more active bots, nothing found")
}
