package adventofcode2016

// Day16 returns checksum.
func Day16(a string, disksize int) string {
	// Pre-allocate buffer for full disk size
	data := make([]byte, disksize)
	copy(data, a)
	n := len(a)

	// Dragon curve expansion
	for n < disksize {
		// New length after expansion: n + 1 + n = 2n + 1
		newLen := 2*n + 1
		if newLen > disksize {
			newLen = disksize
		}

		// Add separator '0' at position n
		if n < disksize {
			data[n] = '0'
		}

		// Add reversed and flipped copy
		for i := 0; i < n && n+1+i < disksize; i++ {
			src := data[n-1-i]
			if src == '0' {
				data[n+1+i] = '1'
			} else {
				data[n+1+i] = '0'
			}
		}
		n = newLen
	}

	// Checksum reduction in-place
	for n%2 == 0 {
		half := n / 2
		for i := 0; i < half; i++ {
			if data[i*2] == data[i*2+1] {
				data[i] = '1'
			} else {
				data[i] = '0'
			}
		}
		n = half
	}

	return string(data[:n])
}

// Reverse reverses a string (kept for compatibility).
func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
