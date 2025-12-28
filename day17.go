package adventofcode2016

import (
	"crypto/md5"
	"encoding/hex"
)

// day17State represents position and path
type day17State struct {
	x, y int
	path []byte
}

// isOpen checks if a door is open based on MD5 nibble (b-f = open)
func isOpen(nibble byte) bool {
	return nibble >= 0xb
}

// md5s returns hex-encoded MD5 hash (for test compatibility)
func md5s(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Day17 returns the shortest (part1) or longest (part2) path through the maze.
func Day17(passcode string, maxMoves int, part1 bool) string {
	passcodeBytes := []byte(passcode)
	passcodeLen := len(passcode)

	// Pre-allocate buffer for MD5 input
	hashInput := make([]byte, passcodeLen+maxMoves)
	copy(hashInput, passcodeBytes)

	var longest string
	queue := []day17State{{0, 3, nil}} // start at (0,3)

	dirs := [4]struct {
		dx, dy int
		ch     byte
	}{
		{0, 1, 'U'},  // up (y increases)
		{0, -1, 'D'}, // down
		{-1, 0, 'L'}, // left
		{1, 0, 'R'},  // right
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if len(cur.path) >= maxMoves {
			continue
		}

		// Build hash input: passcode + path
		copy(hashInput[passcodeLen:], cur.path)
		hash := md5.Sum(hashInput[:passcodeLen+len(cur.path)])

		// Check each direction
		for i, dir := range dirs {
			// Get nibble for this direction (first 4 chars of hex = first 2 bytes)
			var nibble byte
			if i < 2 {
				nibble = hash[0] >> (4 * (1 - i)) & 0xF
			} else {
				nibble = hash[1] >> (4 * (3 - i)) & 0xF
			}

			if !isOpen(nibble) {
				continue
			}

			nx, ny := cur.x+dir.dx, cur.y+dir.dy
			if nx < 0 || nx >= 4 || ny < 0 || ny >= 4 {
				continue
			}

			newPath := make([]byte, len(cur.path)+1)
			copy(newPath, cur.path)
			newPath[len(cur.path)] = dir.ch

			// Check if reached destination (3, 0)
			if nx == 3 && ny == 0 {
				pathStr := string(newPath)
				if part1 {
					return pathStr
				}
				if len(pathStr) > len(longest) {
					longest = pathStr
				}
				continue // Don't add to queue - path ends here
			}

			queue = append(queue, day17State{nx, ny, newPath})
		}
	}

	return longest
}
