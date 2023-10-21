package adventofcode2016

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

const (
	dimX           = 4
	dimY           = 4
	startPosition  = 0 + 3i
	finishPosition = 3 + 0i
)

type direction struct {
	rel complex64 // relative direction such as (0,1) for down
	rep string    // textual representation
	idx int       // index of direction into MD5
}

type state struct {
	pos      complex64
	passcode string
}

func (a state) newState(d direction) state {
	return state{a.pos + d.rel, a.passcode + d.rep}
}

// returns a list of possible moves
func (a state) next() []state {
	var ss []state

	// 4 possible directions in the order of open doors index in MD5
	var (
		up    = direction{0 + 1i, "U", 0}
		down  = direction{0 - 1i, "D", 1}
		left  = direction{-1 + 0i, "L", 2}
		right = direction{1 + 0i, "R", 3}
	)

	// "Any b, c, d, e, or f means that the corresponding door is open"
	open := func(b byte) bool {
		return b >= 'b' && b <= 'f'
	}

	// filter 1: door must be open
	doors := md5s(a.passcode)

	x := real(a.pos)
	y := imag(a.pos)

	if open(doors[up.idx]) && y < dimY-1 {
		ss = append(ss, a.newState(up))
	}
	if open(doors[down.idx]) && y > 0 {
		ss = append(ss, a.newState(down))
	}
	if open(doors[left.idx]) && x > 0 {
		ss = append(ss, a.newState(left))
	}
	if open(doors[right.idx]) && x < dimX-1 {
		ss = append(ss, a.newState(right))
	}

	return ss
}

type states map[state]bool

func (a states) ended() (bool, state) {
	for k := range a {
		if k.pos == finishPosition {
			return true, k
		}
	}
	var zero state
	return false, zero
}

func md5s(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Day17 returns  the shortest path through the maze.
// Cowardly refusing to potentially loop forever
// returns the empty string if maxMoves have been tried and nothing found.
func Day17(passcode string, maxMoves int, part1 bool) string {
	ss := make(states)
	// populate with start position
	ss[state{startPosition, passcode}] = true
	prospects := make(states)

	for i := 0; i < maxMoves; i++ {
		for s := range ss {
			for _, s2 := range s.next() {
				prospects[s2] = true
			}
		}
		finished, winner := ss.ended()
		if finished {
			// strip original passcode from result
			return strings.TrimPrefix(winner.passcode, passcode)
		}

		ss, prospects = prospects, ss
		clear(prospects)
	}
	return ""
}
