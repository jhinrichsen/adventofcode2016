package adventofcode2016

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// Day14 returns 64th key for given salt.
func Day14(salt string, part1 bool) uint {
	const (
		nth  = 64
		next = 1000
	)
	hasher := func() func(string) string {
		if part1 {
			return hash
		}
		return stretchedHash
	}()
	var hashes []string
	keys := uint8(1)
	for i := 0; ; i++ {
		for j := len(hashes); j < i+next+1; j++ {
			hashes = append(hashes, "")
			hashes[j] = hasher(salt + strconv.Itoa(j))
		}
		findTriple := func(s string) byte {
			for i := 0; i < len(s)-2; i++ {
				if s[i] == s[i+1] &&
					s[i] == s[i+2] {
					return s[i]
				}
			}
			return 0
		}
		triple := findTriple(hashes[i])
		if triple == 0 {
			continue
		}
		any := func(idx int, b byte) bool {
			for i := idx; i < idx+next; i++ {
				s := hashes[i]
				for j := 0; j < len(s)-4; j++ {
					if s[j] == b &&
						s[j+1] == b &&
						s[j+2] == b &&
						s[j+3] == b &&
						s[j+4] == b {
						return true
					}
				}
			}
			return false
		}
		if any(i+1, triple) {
			if keys == nth {
				return uint(i)
			}
			keys++
		}
	}
}

func hash(s string) string {
	bs := md5.Sum([]byte(s))
	return hex.EncodeToString(bs[:])
}

func stretchedHash(s string) string {
	buf := []byte(s)
	for i := 0; i < 2017; i++ {
		bs := md5.Sum(buf)
		buf = []byte(hex.EncodeToString(bs[:]))
	}
	return string(buf)
}
