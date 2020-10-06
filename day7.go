package adventofcode2016

import "strings"

// palindrome returns true if s has a four character palindrome ("ABBA").
func palindrome(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		p := s[i] != s[i+1] &&
			s[i+1] == s[i+2] &&
			s[i+3] == s[i]
		if p {
			return true
		}
	}
	return false
}

// Day7 returns how many IP addresses support TLS.
func Day7(lines []string) (n uint) {
	for _, line := range lines {
		var enabled, disabled bool
		inBrackets := true
		var nextBracket byte
		toggle := func() {
			inBrackets = !inBrackets
			if inBrackets {
				nextBracket = ']'
			} else {
				nextBracket = '['
			}
		}
		toggle()
		start := 0
		for start < len(line) {
			// stop is relative to start
			stop := strings.IndexByte(line[start:], nextBracket)
			if stop == -1 {
				stop = len(line) - start
			}
			if palindrome(line[start : start+stop]) {
				if inBrackets {
					disabled = true
					break
				} else {
					enabled = true
					// cannot break yet, need to continue to
					// check for later disables
				}
			}
			start += stop + 1
			toggle()
		}
		if !disabled && enabled {
			n++
		}
	}
	return
}
