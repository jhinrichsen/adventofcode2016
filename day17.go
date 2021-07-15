package adventofcode2016

import (
	"crypto/md5"
	"encoding/hex"
)

type Position complex64

const (
	DimX          = 4
	DimY          = 4
	StartPosition = 0 + 3i
	EndPosition   = 3 + 0i // bottom right row reaches vault
)

type Direction complex64

// Relative directions for one step
const (
	Up    = 1 + 0i
	Down  = -1 + 0i
	Left  = 0 - 1i
	Right = 1 + 0i
)

// Size of the

var Directions = [...]Direction{Up, Down, Left, Right}

func md5lead(buf []byte) string {
	sum := md5.Sum(buf)
	return string(hex.EncodeToString(sum[:][0:2]))
}

// openDoors returns a set of open doors.
func openDoors(hash []byte) map[Direction]bool {
	doors := make(map[Direction]bool)
	for _, b := range hash {
		if b >= 'b' && b <= 'f' {
			idx := b - 'b'
			doors[Directions[idx]] = true
		}
	}
	return doors
}

func wall(p Position) bool {
	x := real(p)
	y := imag(p)
	return x < 0 || x >= DimX || y < 0 || y >= DimY
}
