package adventofcode2016

// numeric returns true if b is an ASCII digit
func numeric(b byte) bool {
	return '0' <= b && b <= '9'
}

// Day04Part1 returns sum of sector IDs of all real rooms.
func Day04Part1(input []byte) uint {
	var sum uint
	var letters [26]byte
	var checksum [5]byte
	var sector uint
	var checksumIdx int
	inChecksum := false

	for i := 0; i < len(input); i++ {
		b := input[i]
		switch {
		case b >= 'a' && b <= 'z':
			if inChecksum {
				if checksumIdx < 5 {
					checksum[checksumIdx] = b
					checksumIdx++
				}
			} else {
				letters[b-'a']++
			}
		case b >= '0' && b <= '9':
			sector = sector*10 + uint(b-'0')
		case b == '[':
			inChecksum = true
		case b == ']':
			// End of checksum
		case b == '\n':
			// End of room - check if real
			if checksumIdx == 5 && isRealRoom(letters, checksum) {
				sum += sector
			}
			// Reset for next room
			letters = [26]byte{}
			checksum = [5]byte{}
			sector = 0
			checksumIdx = 0
			inChecksum = false
		}
	}
	// Handle last room if no trailing newline
	if checksumIdx == 5 && isRealRoom(letters, checksum) {
		sum += sector
	}
	return sum
}

// isRealRoom checks if checksum matches the 5 most common letters (ties broken alphabetically)
func isRealRoom(letters [26]byte, checksum [5]byte) bool {
	// Find top 5 letters by frequency, alphabetical tiebreaker
	var result [5]byte
	var used [26]bool

	for pos := 0; pos < 5; pos++ {
		bestIdx := -1
		var bestCount byte
		for i := 0; i < 26; i++ {
			if used[i] {
				continue
			}
			if letters[i] > bestCount {
				bestCount = letters[i]
				bestIdx = i
			}
		}
		if bestIdx == -1 || bestCount == 0 {
			return false
		}
		result[pos] = byte(bestIdx) + 'a'
		used[bestIdx] = true
	}
	return result == checksum
}

// Day04Part2 returns sector ID of decrypted real room "northpole object storage".
func Day04Part2(input []byte) uint {
	// Target decrypted: "northpole object storage"
	// Encrypted format has dashes at positions 9, 16, 24 (before sector)
	// Quick filter: check dash positions before full parsing

	var roomStart int
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			room := input[roomStart:i]
			// Quick filter: "northpole object storage" encrypted has dashes at 9, 16, 24
			if len(room) > 25 && room[9] == '-' && room[16] == '-' && room[24] == '-' {
				// Parse and check this room
				var letters [26]byte
				var checksum [5]byte
				var sector uint
				var checksumIdx int
				inChecksum := false

				for j := 0; j < len(room); j++ {
					b := room[j]
					switch {
					case b >= 'a' && b <= 'z':
						if inChecksum {
							if checksumIdx < 5 {
								checksum[checksumIdx] = b
								checksumIdx++
							}
						} else {
							letters[b-'a']++
						}
					case b >= '0' && b <= '9':
						sector = sector*10 + uint(b-'0')
					case b == '[':
						inChecksum = true
					}
				}

				if checksumIdx == 5 && isRealRoom(letters, checksum) {
					if decryptsToTarget(room, sector) {
						return sector
					}
				}
			}
			roomStart = i + 1
		}
	}
	return 0
}

// decryptsToTarget checks if room decrypts to "northpole object storage"
func decryptsToTarget(room []byte, sector uint) bool {
	target := []byte("northpole object storage")
	n := byte(sector % 26)
	targetIdx := 0

	for i := 0; i < len(room); i++ {
		b := room[i]
		if b >= '0' && b <= '9' {
			break
		}
		if b == '-' {
			// Final dash before sector means we're done with the name
			if targetIdx >= len(target) {
				break
			}
			if target[targetIdx] != ' ' {
				return false
			}
			targetIdx++
			continue
		}
		if b >= 'a' && b <= 'z' {
			if targetIdx >= len(target) {
				return false
			}
			decrypted := b + n
			if decrypted > 'z' {
				decrypted -= 26
			}
			if target[targetIdx] != decrypted {
				return false
			}
			targetIdx++
		}
	}
	return targetIdx == len(target)
}
