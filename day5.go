package adventofcode2016

import (
	"crypto/md5"
	"fmt"
)

const day5Input = "uqwqemis"

// Day5 solves day 1 part 1.
func Day5(doorID string, part1 bool) string {
	var password [8]byte
	var idx int // 0..n-1
	for i := 0; ; i++ {
		chk := md5.Sum([]byte(fmt.Sprintf("%s%d", doorID, i)))
		fiveZeroes := chk[0] == 0 &&
			chk[1] == 0 &&
			chk[2] < 16
		if !fiveZeroes {
			continue
		}
		rep := []byte(fmt.Sprintf("%x", chk))
		if part1 {
			// next character is at position 6
			password[idx] = rep[5]
		} else {
			// next character is index of next next
			// character
			j := int(rep[5] - '0')
			if j >= len(password) {
				continue // skip illegal index, try next
			}
			// "Use only the first result for each position"
			if password[j] != 0 {
				continue
			}
			password[j] = rep[6]
		}
		idx++
		if idx == len(password) {
			break
		}
	}
	return string(password[:])
}

// Day5Part2 solves day 1 part 1.
func Day5Part2(lines []string) uint {
	return 0
}
