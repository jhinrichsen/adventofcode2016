package adventofcode2016

// Day9Part1 returns decompressed length of strings.
func Day9Part1(s string) (sum uint) {
	var times uint = 1        // how many times to repeat a character
	var repeat uint           // number of characters to repeat
	var inMarker bool         // state maching: parsing marker
	var marker1, marker2 uint // temporary marker content while parsing

	// one sweet single pass
	for _, b := range s {
		if b == '(' && repeat == 0 { // enter marker, no other marker active
			inMarker = true
		} else if b == ')' && repeat == 0 { // exit marker
			inMarker = false
			repeat, times = marker1, marker2
			marker1, marker2 = 0, 0
		} else if inMarker {
			if b == 'x' { // marker separator?
				marker1 = marker2
				marker2 = 0
			} else {
				marker2 *= 10
				marker2 += uint(b - '0')
			}
		} else {
			sum += times // add 1, or x * 1

			// chars wear off one by one in active zone
			if repeat > 0 {
				repeat--
			}

			// if there are no more chars to repeat, reset times
			if repeat == 0 {
				times = 1
			}
		}
	}
	return
}

// Day9Part2 returns decompressed length of strings, version 2.
// Linear iteration from left to right won't cut it this time, because multiple
// repeats need to be considered, probably needing a stack container e.a.
// Instead, iterate from right to left, resolving any (repeat x time)
// immediately.
func Day9Part2(s string) (sum uint) {
	var inMarker bool         // state maching: parsing marker
	var marker1, marker2 uint // temporary marker content while parsing
	factors := make([]uint, len(s)+1)
	for i := 0; i < len(factors); i++ {
		factors[i] = 1
	}

	// one sweet single pass, this time from right to left
	for i, b := range s {
		if b == '(' { // enter marker
			inMarker = true
		} else if b == ')' { // exit marker
			inMarker = false

			// set factor = times * factor for index [i..i+repeat]
			start := uint(i) + 1
			stop := start + marker1
			for j := start; j < stop; j++ {
				factors[j] *= marker2
			}
			marker1, marker2 = 0, 0
		} else if inMarker {
			if b == 'x' { // marker separator?
				marker1 = marker2
				marker2 = 0
			} else {
				marker2 *= 10
				marker2 += uint(b - '0')
			}
		} else {
			sum += factors[i]
		}
	}
	return
}
