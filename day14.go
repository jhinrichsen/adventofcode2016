package adventofcode2016

import (
	"crypto/md5"
)

// hexTable for converting nibbles to hex chars
var hexTable = [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

// Day14 returns 64th key for given salt.
func Day14(salt string, part1 bool) uint {
	const (
		nth  = 64
		next = 1000
	)

	// Pre-allocate buffer for salt + number
	saltBytes := []byte(salt)
	buf := make([]byte, len(salt)+20)
	copy(buf, saltBytes)
	base := len(salt)

	// Hash cache
	var hashes [][32]byte // store raw hash bytes, not hex strings

	// Compute hash for index, using stretched version if part2
	computeHash := func(idx int) [16]byte {
		n := writeInt(buf[base:], idx)
		if part1 {
			return md5.Sum(buf[:base+n])
		}
		return stretchedHashBytes(buf[:base+n])
	}

	// Convert hash to hex string (32 bytes)
	hashToHex := func(h [16]byte) [32]byte {
		var hex [32]byte
		for i, b := range h {
			hex[i*2] = hexTable[b>>4]
			hex[i*2+1] = hexTable[b&0x0f]
		}
		return hex
	}

	keys := uint8(1)
	for i := 0; ; i++ {
		// Ensure we have enough hashes computed
		for j := len(hashes); j < i+next+1; j++ {
			h := computeHash(j)
			hashes = append(hashes, hashToHex(h))
		}

		// Find triple in hash[i]
		hex := hashes[i]
		triple := byte(0)
		for j := 0; j < 30; j++ {
			if hex[j] == hex[j+1] && hex[j] == hex[j+2] {
				triple = hex[j]
				break
			}
		}
		if triple == 0 {
			continue
		}

		// Look for quint in next 1000 hashes
		found := false
		for k := i + 1; k <= i+next; k++ {
			hex := hashes[k]
			for j := 0; j < 28; j++ {
				if hex[j] == triple &&
					hex[j+1] == triple &&
					hex[j+2] == triple &&
					hex[j+3] == triple &&
					hex[j+4] == triple {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if found {
			if keys == nth {
				return uint(i)
			}
			keys++
		}
	}
}

// stretchedHashBytes computes MD5 2017 times, reusing buffers
func stretchedHashBytes(input []byte) [16]byte {
	var hexBuf [32]byte
	h := md5.Sum(input)

	for i := 0; i < 2016; i++ {
		// Convert hash to hex
		for j, b := range h {
			hexBuf[j*2] = hexTable[b>>4]
			hexBuf[j*2+1] = hexTable[b&0x0f]
		}
		h = md5.Sum(hexBuf[:])
	}
	return h
}

// stretchedHash returns stretched hash as hex string (for test compatibility)
func stretchedHash(s string) string {
	h := stretchedHashBytes([]byte(s))
	var hex [32]byte
	for i, b := range h {
		hex[i*2] = hexTable[b>>4]
		hex[i*2+1] = hexTable[b&0x0f]
	}
	return string(hex[:])
}
