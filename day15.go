package adventofcode2016

import (
	"fmt"
	"strconv"
	"strings"
)

type day15 []disc

type disc struct {
	position  int // position at time = 0
	positions int // number of positions
}

func (a disc) tick(time int) int {
	return (a.position + time) % a.positions
}

// newDay15 parses ascending lines in the form "Disc #1 has 5 positions; at
// time=0, it is at position 4.".
func newDay15(lines []string) (day15, error) {
	var d day15
	for i, line := range lines {
		fs := strings.Fields(line)
		if len(fs) != 12 {
			return d, fmt.Errorf("line %d: want 15 columns but got %q",
				i, line)
		}
		positions, err := strconv.Atoi(fs[3])
		if err != nil {
			return d, fmt.Errorf("cannot parse number %q", fs[3])
		}
		position, err := strconv.Atoi(fs[6][4:])
		if err != nil {
			return d, err
		}
		d = append(d, disc{
			position:  position,
			positions: positions,
		})
	}
	return d, nil
}

// Day15 returns time when cylinder drops through.
func Day15(d day15, part1 bool) uint {
	return Day15Hardcoded(d, part1)
}

// Day15Broken returns time when a capsule passes through all discs.
func Day15Broken(d day15, part1 bool) uint {
	p1 := -1
	deepest := -1

	var done bool
	var time int
	for step := 1; !done; time += step {
		done = true
		for i := range d {
			fallThrough := d[i].tick(time+1+i) == 0
			if fallThrough && i > deepest {
				step *= d[i].positions
				deepest = i
			}
			if !fallThrough {
				done = false
				if p1 == -1 && i == len(d)-1 {
					p1 = time // part 1
				}
				break
			}
		}
	}
	if part1 {
		return uint(p1)
	}
	return uint(time)
}

// Day15Hardcoded features a mostly hardcoded dead simple no frills
// implementation.
func Day15Hardcoded(d day15, part1 bool) uint {
	d1 := func(t uint) bool { return (t+1)%17 == 0 }
	d2 := func(t uint) bool { return (t+0)%7 == 0 }
	d3 := func(t uint) bool { return (t+2)%19 == 0 }
	d4 := func(t uint) bool { return (t+0)%5 == 0 }
	d5 := func(t uint) bool { return (t+0)%3 == 0 }
	d6 := func(t uint) bool { return (t+5)%13 == 0 }

	d7 := func(t uint) bool { return (t+0)%11 == 0 }

	for t := uint(0); ; t++ {
		b := d1(t+1) &&
			d2(t+2) &&
			d3(t+3) &&
			d4(t+4) &&
			d5(t+5) &&
			d6(t+6)
		if !part1 {
			b = b && d7(t+7)
		}
		if b {
			return t
		}
	}
}

// Day15Smart solves day 15.
// Some day, i will come back and implement a full blown chinese remainder
// theorem. I _really_ need to work on my math background, Gauss 1801,
// Disquisitiones Arithmeticae. Modular multiplicative inversions, pairwise
// coprimes, ...
// Disk sizes are small enough though, says askalski.
func Day15Smart(d day15, part1 bool) uint {
	t := 0
	inc := 1
	for i, disc := range d {
		for (t+disc.position+(i+1))%disc.positions != 0 {
			t += inc
		}
		inc *= disc.positions
	}
	return uint(t)
}
