package adventofcode2016

import (
	"crypto/md5"
)

// Day05 solves day 5.
func Day05(doorID string, part1 bool) string {
	var password [8]byte
	var found int

	// Pre-allocate buffer for doorID + number
	// Max int is ~10 digits, so 20 bytes should be plenty
	buf := make([]byte, len(doorID)+20)
	copy(buf, doorID)
	base := len(doorID)

	for i := 0; found < 8; i++ {
		// Write number to buffer
		n := writeInt(buf[base:], i)

		chk := md5.Sum(buf[:base+n])

		// Check for 5 leading zeros (first 20 bits = 0)
		if chk[0] != 0 || chk[1] != 0 || chk[2] >= 16 {
			continue
		}

		// 6th hex char is low nibble of chk[2]
		sixth := chk[2] & 0x0f
		// 7th hex char is high nibble of chk[3]
		seventh := (chk[3] >> 4) & 0x0f

		if part1 {
			password[found] = hexChar(sixth)
			found++
		} else {
			// 6th char is position (0-7)
			pos := int(sixth)
			if pos >= 8 || password[pos] != 0 {
				continue
			}
			password[pos] = hexChar(seventh)
			found++
		}
	}
	return string(password[:])
}

// writeInt writes a non-negative integer to buf and returns the number of bytes written
func writeInt(buf []byte, n int) int {
	if n == 0 {
		buf[0] = '0'
		return 1
	}

	// Count digits
	temp := n
	digits := 0
	for temp > 0 {
		digits++
		temp /= 10
	}

	// Write digits in reverse
	for i := digits - 1; i >= 0; i-- {
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return digits
}

// hexChar returns the hex character for a value 0-15
func hexChar(v byte) byte {
	if v < 10 {
		return '0' + v
	}
	return 'a' + v - 10
}

// Day5Part2 solves day 5 part 2 (unused, kept for compatibility).
func Day5Part2(lines []string) uint {
	return 0
}
