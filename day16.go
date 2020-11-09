package adventofcode2016

import "strings"

// Day16 returns checksum.
func Day16(a string, disksize int) string {
	for len(a) < disksize {
		b := []byte(Reverse(a))

		// flip
		for i := range b {
			if b[i] == '0' {
				b[i] = '1'
			} else {
				b[i] = '0'
			}
		}
		a = a + "0" + string(b)
	}

	// fit to disk
	a = a[:disksize]

	odd := func(s string) bool {
		return len(s)%2 != 0
	}

	// reduce
	for {
		var sb strings.Builder
		for pair := 0; pair < len(a); pair += 2 {
			if a[pair] == a[pair+1] {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}

		a = sb.String()
		if odd(a) {
			break
		}
	}
	return a
}

// Reverse reverses a string.
// Implemenation taken from
// https://github.com/golang/example/blob/master/stringutil/reverse.go
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
