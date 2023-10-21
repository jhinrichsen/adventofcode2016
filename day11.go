package adventofcode2016

import (
	"fmt"
)

type day11 struct {
	fl1, fl2, fl3, fl4 floor
	floorID            int
	current            *floor
	distance           int
}

type floor struct {
	chips, generators int
}

// Day11 solves day 11.
func Day11(part1 bool) (uint, error) {
	var items int
	if part1 {
		items = 5
	} else {
		items = 7
	}

	visited := make(map[string]int)
	floor1 := floor{items - 2, items} // In my input state, 2 microchips were on floor 2
	c := make(chan day11, 100000)
	c <- day11{floor1, floor{2, 0}, floor{}, floor{}, 1, &floor1, 0}
	for state := range c {
		for _, route := range state.adjacent() {
			if route.isSolved() {
				return uint(state.distance + 1), nil

			}
			if _, seen := visited[route.String()]; !seen {
				route.distance = state.distance + 1
				visited[route.String()] = route.distance
				c <- route
			}
		}
	}
	return 0, fmt.Errorf("unsolvable")
}

func cap2(n int) int {
	if n > 2 {
		return 2
	}
	return n
}

// Adjacent states from valid moves from this state
func (s day11) adjacent() (routes []day11) {
	maxG, maxM := cap2(s.current.generators), cap2(s.current.chips) //Only bring up to 2 items up/down
	if s.floorID != 4 {                                             //Move items up in all combinations
		for g := 0; g <= maxG; g++ {
			for m := 0; m <= maxM; m++ {
				if !(g+m == 0 || g+m > 2 || (maxM+maxG >= 2 && g+m == 1)) {
					routes = append(routes, getNewState(s, true, m, g))
				}
			}
		}
	}
	if s.floorID != 1 { // Move items down in all combinations
		for g := 0; g <= maxG; g++ {
			for m := 0; m <= maxM; m++ {
				if !(g+m == 0 || g+m > 1) {
					routes = append(routes, getNewState(s, false, m, g))
				}
			}
		}
	}
	return
}

func (s day11) isSolved() bool {
	return s.fl1.isEmpty() && s.fl2.isEmpty() && s.fl3.isEmpty()
}
func (s day11) String() string {
	return fmt.Sprintf("%s%s%s%s%v", s.fl1, s.fl2, s.fl3, s.fl4, s.floorID)
}

func (fl floor) isEmpty() bool {
	return fl.chips == 0 && fl.generators == 0
}
func (fl floor) String() string {
	return fmt.Sprintf("%v%v", fl.chips, fl.generators)
}

func getNewState(s day11, up bool, c, g int) day11 {
	switch s.floorID {
	case 1:
		newFloor := floor{s.fl2.chips + c, s.fl2.generators + g}
		return day11{floor{s.fl1.chips - c, s.fl1.generators - g}, newFloor, s.fl3, s.fl4, 2, &newFloor, 0}
	case 2:
		if up {
			newFloor := floor{s.fl3.chips + c, s.fl3.generators + g}
			return day11{s.fl1, floor{s.fl2.chips - c, s.fl2.generators - g}, newFloor, s.fl4, 3, &newFloor, 0}
		}
		newFloor := floor{s.fl1.chips + c, s.fl1.generators + g}
		return day11{newFloor, floor{s.fl2.chips - c, s.fl2.generators - g}, s.fl3, s.fl4, 1, &newFloor, 0}
	case 3:
		if up {
			newFloor := floor{s.fl4.chips + c, s.fl4.generators + g}
			return day11{s.fl1, s.fl2, floor{s.fl3.chips - c, s.fl3.generators - g}, newFloor, 4, &newFloor, 0}
		}
		newFloor := floor{s.fl2.chips + c, s.fl2.generators + g}
		return day11{s.fl1, newFloor, floor{s.fl3.chips - c, s.fl3.generators - g}, s.fl4, 2, &newFloor, 0}
	case 4:
		newFloor := floor{s.fl3.chips + c, s.fl3.generators + g}
		return day11{s.fl1, s.fl2, newFloor, floor{s.fl1.chips - c, s.fl1.generators - g}, 3, &newFloor, 0}
	}
	panic("Invalid state")
}
