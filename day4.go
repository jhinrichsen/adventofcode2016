package adventofcode2016

import (
	"fmt"
	"sort"
	"strings"
)

type day4 struct {
	ID       []byte
	letters  map[byte]uint
	sector   uint
	checksum [5]byte
}

// avoid rune conversion
func numeric(b byte) bool {
	return '0' <= b && b <= '9'
}

const dash = '-'

func newDay4(s string) (day4, error) {
	var d day4
	d.ID = []byte(s)
	d.letters = make(map[byte]uint, 26)

	// sparse map is faster than pre-allocating all letters to 0
	/*
		for i := byte('a'); i <= 'z'; i++ {
			d.letters[i] = 0
		}
	*/
	letter := func(b byte) {
		// insert into sparse map
		if _, ok := d.letters[b]; ok {
			d.letters[b]++
		} else {
			d.letters[b] = 1
		}
	}
	for i, b := range d.ID {
		if b == dash {
			continue
		}
		if numeric(b) {
			d.sector = d.sector*10 + uint(b) - '0'
			continue
		}
		if b == '[' {
			for j := 0; j < len(d.checksum); j++ {
				d.checksum[j] = d.ID[i+1+j]
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

// Day4Part1 returns sum of sector IDs of all real rooms.
func Day4Part1(lines []string) (uint, error) {
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

// Day4Part2 returns sector ID of decrypted real room "northpole object
// storage".
func Day4Part2(lines []string) (uint, error) {
	const room = "northpole object storage"
	for i, line := range lines {
		// has word separator at the right place?
		p1 := line[9] == dash &&
			line[16] == dash &&
			line[24] == dash
		if !p1 {
			continue
		}
		d, err := newDay4(line)
		if err != nil {
			return 0, fmt.Errorf("error in line %d: %w", i+1, err)
		}
		p2 := d.real()
		if !p2 {
			continue
		}
		if decrypt(line) == room {
			return d.sector, nil
		}
	}
	return 0, fmt.Errorf("not found")
}

func decrypt(room string) string {
	d, _ := newDay4(room)
	n := d.sector % 26
	var sb strings.Builder
	for i := 0; i < len(room); i++ {
		b := room[i]
		if numeric(b) {
			break
		}
		if b == '-' {
			sb.WriteByte(' ')
			continue
		}
		b += byte(n)
		if b > 'z' {
			b -= 26
		}
		sb.WriteByte(b)
	}
	// strip trailing space
	return strings.TrimSpace(sb.String())
}
