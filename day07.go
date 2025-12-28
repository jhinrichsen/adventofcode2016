package adventofcode2016

// Day07 returns how many IP addresses support TLS (part1=true) or SSL (part1=false).
func Day07(input []byte, part1 bool) uint {
	var count uint
	lineStart := 0

	for i := 0; i <= len(input); i++ {
		if i == len(input) || input[i] == '\n' {
			if i > lineStart {
				line := input[lineStart:i]
				if part1 {
					if supportsTLS(line) {
						count++
					}
				} else {
					if supportsSSL(line) {
						count++
					}
				}
			}
			lineStart = i + 1
		}
	}
	return count
}

// supportsTLS checks if an IP supports TLS (has ABBA outside brackets, none inside)
func supportsTLS(line []byte) bool {
	hasABBAOutside := false
	hasABBAInside := false
	inBrackets := false

	for i := 0; i < len(line)-3; i++ {
		b := line[i]
		if b == '[' {
			inBrackets = true
			continue
		}
		if b == ']' {
			inBrackets = false
			continue
		}
		// Check for ABBA pattern: s[i] != s[i+1] && s[i+1] == s[i+2] && s[i+3] == s[i]
		if line[i] != line[i+1] && line[i+1] == line[i+2] && line[i+3] == line[i] &&
			line[i+1] != '[' && line[i+1] != ']' && line[i+2] != '[' && line[i+2] != ']' && line[i+3] != '[' && line[i+3] != ']' {
			if inBrackets {
				hasABBAInside = true
			} else {
				hasABBAOutside = true
			}
		}
	}
	return hasABBAOutside && !hasABBAInside
}

// supportsSSL checks if an IP supports SSL (has ABA outside and corresponding BAB inside)
func supportsSSL(line []byte) bool {
	// Collect all ABA patterns outside brackets, check for matching BAB inside
	// Use a simple approach: for each ABA outside, scan inside for BAB

	// First pass: find all bracket regions
	type region struct {
		start, end int
		inside     bool
	}
	var regions []region
	regionStart := 0
	inside := false
	for i := 0; i < len(line); i++ {
		if line[i] == '[' {
			if i > regionStart {
				regions = append(regions, region{regionStart, i, inside})
			}
			inside = true
			regionStart = i + 1
		} else if line[i] == ']' {
			if i > regionStart {
				regions = append(regions, region{regionStart, i, inside})
			}
			inside = false
			regionStart = i + 1
		}
	}
	if len(line) > regionStart {
		regions = append(regions, region{regionStart, len(line), inside})
	}

	// Find ABA patterns outside brackets
	for _, r := range regions {
		if r.inside {
			continue
		}
		for i := r.start; i < r.end-2; i++ {
			// ABA pattern: line[i] != line[i+1] && line[i+2] == line[i]
			if line[i] != line[i+1] && line[i+2] == line[i] {
				a, b := line[i], line[i+1]
				// Look for BAB inside brackets
				for _, r2 := range regions {
					if !r2.inside {
						continue
					}
					for j := r2.start; j < r2.end-2; j++ {
						if line[j] == b && line[j+1] == a && line[j+2] == b {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
