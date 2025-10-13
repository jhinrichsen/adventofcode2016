package adventofcode2016

import (
	"image"
	"math"
)

type Day24Puzzle struct {
	grid      [][]byte
	dimY      int
	dimX      int
	locations map[byte]image.Point
	numLocs   int
}

func NewDay24(lines []string) Day24Puzzle {
	dimY := len(lines)
	grid := make([][]byte, dimY)
	locations := make(map[byte]image.Point)

	for y := range grid {
		grid[y] = []byte(lines[y])
		for x, cell := range grid[y] {
			if cell >= '0' && cell <= '9' {
				locations[cell] = image.Point{X: x, Y: y}
			}
		}
	}

	return Day24Puzzle{
		grid:      grid,
		dimY:      dimY,
		dimX:      len(lines[0]),
		locations: locations,
		numLocs:   len(locations),
	}
}

func Day24(puzzle Day24Puzzle, part1 bool) uint {
	// Calculate distances between all pairs of locations using BFS
	distances := make(map[[2]byte]uint)

	for from := range puzzle.locations {
		dists := bfsDistances(puzzle, from)
		for to, dist := range dists {
			distances[[2]byte{from, to}] = dist
		}
	}

	// Use dynamic programming with bitmasks to solve TSP
	// State: (current location, visited mask)
	memo := make(map[[2]uint]uint)
	allVisited := (1 << puzzle.numLocs) - 1
	startMask := uint(1) // location '0' is visited

	var tsp func(current byte, visited uint) uint
	tsp = func(current byte, visited uint) uint {
		if visited == uint(allVisited) {
			if part1 {
				return 0 // Part 1: don't need to return to start
			}
			// Part 2: return to location '0'
			return distances[[2]byte{current, '0'}]
		}

		state := [2]uint{uint(current), visited}
		if result, exists := memo[state]; exists {
			return result
		}

		minDist := uint(math.MaxUint)

		for target := range puzzle.locations {
			bit := uint(1) << (target - '0')
			if visited&bit == 0 { // not visited yet
				newVisited := visited | bit
				dist := distances[[2]byte{current, target}] + tsp(target, newVisited)
				if dist < minDist {
					minDist = dist
				}
			}
		}

		memo[state] = minDist
		return minDist
	}

	return tsp('0', startMask)
}

func bfsDistances(puzzle Day24Puzzle, start byte) map[byte]uint {
	startPos := puzzle.locations[start]
	distances := make(map[byte]uint)
	visited := make([][]bool, puzzle.dimY)
	for y := range visited {
		visited[y] = make([]bool, puzzle.dimX)
	}

	queue := []image.Point{startPos}
	visited[startPos.Y][startPos.X] = true
	dist := uint(0)

	directions := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // up, right, down, left

	for len(queue) > 0 {
		size := len(queue)

		for range size {
			current := queue[0]
			queue = queue[1:]

			// Check if this position contains a numbered location
			cell := puzzle.grid[current.Y][current.X]
			if cell >= '0' && cell <= '9' {
				distances[cell] = dist
			}

			// Explore neighbors
			for _, dir := range directions {
				next := image.Point{X: current.X + dir.X, Y: current.Y + dir.Y}

				if next.X >= 0 && next.X < puzzle.dimX &&
				   next.Y >= 0 && next.Y < puzzle.dimY &&
				   !visited[next.Y][next.X] &&
				   puzzle.grid[next.Y][next.X] != '#' {
					visited[next.Y][next.X] = true
					queue = append(queue, next)
				}
			}
		}
		dist++
	}

	return distances
}
