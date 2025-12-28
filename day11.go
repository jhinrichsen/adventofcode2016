package adventofcode2016

// day11State packs the state into uint32:
// bits 0-3: floor1 chips, 4-7: floor1 gens
// bits 8-11: floor2 chips, 12-15: floor2 gens
// bits 16-19: floor3 chips, 20-23: floor3 gens
// bits 24-27: floor4 chips, 28-31: floor4 gens
// We store current floor separately
type day11State struct {
	floors   uint32
	floorID  uint8
	distance uint16
}

func (s day11State) getFloor(f int) (chips, gens int) {
	shift := uint(f * 8)
	chips = int((s.floors >> shift) & 0xF)
	gens = int((s.floors >> (shift + 4)) & 0xF)
	return
}

func (s *day11State) setFloor(f, chips, gens int) {
	shift := uint(f * 8)
	mask := uint32(0xFF) << shift
	s.floors = (s.floors &^ mask) | (uint32(chips&0xF) << shift) | (uint32(gens&0xF) << (shift + 4))
}

func (s day11State) key() uint64 {
	return uint64(s.floors) | (uint64(s.floorID) << 32)
}

func (s day11State) isSolved() bool {
	// All items on floor 4 means floors 1-3 are empty
	return s.floors&0x00FFFFFF == 0
}

// Day11 solves day 11.
func Day11(part1 bool) (uint, error) {
	var items int
	if part1 {
		items = 5
	} else {
		items = 7
	}

	// Initial state: floor1 has (items-2) chips and items generators
	// floor2 has 2 chips and 0 generators
	var initial day11State
	initial.setFloor(0, items-2, items) // floor 1
	initial.setFloor(1, 2, 0)           // floor 2
	initial.floorID = 1

	visited := make(map[uint64]bool)
	visited[initial.key()] = true

	queue := []day11State{initial}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		chips, gens := state.getFloor(int(state.floorID) - 1)
		maxM := min(chips, 2)
		maxG := min(gens, 2)

		// Try moving up
		if state.floorID < 4 {
			for g := 0; g <= maxG; g++ {
				for m := 0; m <= maxM; m++ {
					total := g + m
					if total == 0 || total > 2 {
						continue
					}
					// Optimization: prefer moving 2 items up
					if maxM+maxG >= 2 && total == 1 {
						continue
					}
					newState := state.move(true, m, g)
					if newState.isSolved() {
						return uint(state.distance + 1), nil
					}
					key := newState.key()
					if !visited[key] {
						visited[key] = true
						newState.distance = state.distance + 1
						queue = append(queue, newState)
					}
				}
			}
		}

		// Try moving down
		if state.floorID > 1 {
			for g := 0; g <= maxG; g++ {
				for m := 0; m <= maxM; m++ {
					total := g + m
					if total == 0 || total > 1 {
						continue
					}
					newState := state.move(false, m, g)
					key := newState.key()
					if !visited[key] {
						visited[key] = true
						newState.distance = state.distance + 1
						queue = append(queue, newState)
					}
				}
			}
		}
	}
	return 0, nil
}

func (s day11State) move(up bool, chips, gens int) day11State {
	newState := s
	srcFloor := int(s.floorID) - 1
	var dstFloor int
	if up {
		dstFloor = srcFloor + 1
		newState.floorID++
	} else {
		dstFloor = srcFloor - 1
		newState.floorID--
	}

	srcC, srcG := s.getFloor(srcFloor)
	dstC, dstG := s.getFloor(dstFloor)

	newState.setFloor(srcFloor, srcC-chips, srcG-gens)
	newState.setFloor(dstFloor, dstC+chips, dstG+gens)

	return newState
}

func (s day11State) isValid() bool {
	// A floor is safe if:
	// - No generators present (chips are safe), OR
	// - No chips present, OR
	// - All chips have matching generators (chips == gens when treating pairs as equivalent)
	// Since we count pairs, chips on a floor with generators means chips must equal gens
	for f := 0; f < 4; f++ {
		chips, gens := s.getFloor(f)
		// Invalid if there are chips AND generators AND they're not all paired
		if chips > 0 && gens > 0 && chips != gens {
			return false
		}
	}
	return true
}
