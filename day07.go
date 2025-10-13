package adventofcode2016

import (
	"strings"
)

// palindrome3 returns true if s has a n=3 character palindrome ("xyx").
func palindromes3(s string) map[string]bool {
	ps := make(map[string]bool)
	for i := 0; i < len(s)-2; i++ {
		p := s[i] != s[i+1] &&
			s[i+2] == s[i]
		if p {
			ps[s[i:i+3]] = true
		}
	}
	return ps
}

// palindrome4 returns true if s has a n=4 character palindrome ("abba").
func palindromes4(s string) map[string]bool {
	ps := make(map[string]bool)
	for i := 0; i < len(s)-3; i++ {
		p := s[i] != s[i+1] &&
			s[i+1] == s[i+2] &&
			s[i+3] == s[i]
		if p {
			ps[s[i:i+4]] = true
		}
	}
	return ps
}

// Day7 returns how many IP addresses support TLS (part1=true) or SSL
// (part1=false).
func Day07(lines []string, part1 bool) (n uint) {
	var bracketPalindromes map[string]bool
	var otherPalindromes map[string]bool

	var palindromes func(string) map[string]bool
	if part1 {
		palindromes = palindromes4
	} else {
		palindromes = palindromes3
	}

	part1Pred := func() bool {
		disabled := len(bracketPalindromes) > 0
		enabled := len(otherPalindromes) > 0
		return !disabled && enabled
	}
	hasCorresponding := func(s string) bool {
		var rev strings.Builder
		rev.WriteByte(s[1])
		rev.WriteByte(s[0])
		rev.WriteByte(s[1])
		return otherPalindromes[rev.String()]
	}
	part2Pred := func() bool {
		// find any
		for k := range bracketPalindromes {
			if hasCorresponding(k) {
				return true
			}
		}
		return false
	}
	pred := func() func() bool {
		if part1 {
			return part1Pred
		}
		return part2Pred
	}()
	for _, line := range lines {
		bracketPalindromes = make(map[string]bool)
		otherPalindromes = make(map[string]bool)
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
			ps := palindromes(line[start : start+stop])
			if inBrackets {
				for k := range ps {
					bracketPalindromes[k] = true
				}
			} else {
				for k := range ps {
					otherPalindromes[k] = true
				}
			}
			start += stop + 1
			toggle()
		}
		if pred() {
			n++
		}
	}
	return
}
