package adventofcode2016

import (
	"fmt"
	"sort"
)

type day4 struct {
	letters  map[byte]uint
	sector   uint
	checksum [5]byte
}

func newDay4(s string) (day4, error) {
	var d day4
	d.letters = make(map[byte]uint, len(s)) // a bit too large = large enough
	letter := func(b byte) {
		if _, ok := d.letters[b]; ok {
			d.letters[b]++
		} else {
			d.letters[b] = 1
		}
	}
	// avoid rune conversion
	numeric := func(b byte) bool {
		return '0' <= b && b <= '9'
	}
	buf := []byte(s)
	for i, b := range buf {
		if b == '-' {
			continue
		}
		if numeric(b) {
			d.sector = d.sector*10 + uint(b) - '0'
			continue
		}
		if b == '[' {
			for j := 0; j < len(d.checksum); j++ {
				d.checksum[j] = buf[i+1+j]
			}
			return d, nil
		}
		letter(b)
	}
	return d, fmt.Errorf("missing checksum in %q", s)
}

type day4Sector struct {
	b byte
	n uint
}

func (d day4) real() bool {
	// sort letters by occurence
	var sectors []day4Sector
	for k, v := range d.letters {
		sectors = append(sectors, day4Sector{k, v})
	}
	sort.Slice(sectors, func(i, j int) bool {
		if sectors[i].n < sectors[j].n {
			return false
		}
		if sectors[i].n > sectors[j].n {
			return true
		}
		return sectors[i].b < sectors[j].b
	})
	var ck [5]byte
	for i := 0; i < len(ck); i++ {
		ck[i] = sectors[i].b
	}
	return ck == d.checksum
}

// Day4 returns sum of sector IDs of all real rooms.
func Day4(lines []string) (uint, error) {
	var sum uint
	for i, line := range lines {
		d, err := newDay4(line)
		if err != nil {
			return 0, fmt.Errorf("error in line %d: %w", i+1, err)
		}
		if d.real() {
			sum += d.sector
		}
	}
	return sum, nil
}
