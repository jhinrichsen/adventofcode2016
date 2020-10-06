package adventofcode2016

import "strings"

// Day6 returns error corrected transmission.
func Day6(lines []string, part1 bool) string {
	var sb strings.Builder
	for x := 0; x < len(lines[0]); x++ {
		var freqs [26]uint
		for y := 0; y < len(lines); y++ {
			c := lines[y][x]
			freqs[c-'a']++
		}
		var idx byte
		var val uint
		if !part1 {
			val = uint(len(lines)) // initialize proper minimum
		}
		for i, freq := range freqs {
			if part1 {
				if freq > val {
					idx, val = byte(i), freq
				}
			} else if !part1 {
				if 0 < freq && freq < val {
					idx, val = byte(i), freq
				}
			}
		}
		sb.WriteByte(idx + 'a')
	}
	return sb.String()
}
