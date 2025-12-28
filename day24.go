package adventofcode2016

import (
	"image"
	"math"
)

type Day24Puzzle struct {
	grid      [][]byte
	dimY      int
	dimX      int
	locations [10]image.Point // locations 0-9
	numLocs   int
}

func NewDay24(lines []string) Day24Puzzle {
	dimY := len(lines)
	grid := make([][]byte, dimY)
	var locations [10]image.Point
	numLocs := 0

	for y := range grid {
		grid[y] = []byte(lines[y])
		for x, cell := range grid[y] {
			if cell >= '0' && cell <= '9' {
				idx := int(cell - '0')
				locations[idx] = image.Point{X: x, Y: y}
				if idx >= numLocs {
					numLocs = idx + 1
				}
			}
		}
	}

	return Day24Puzzle{
		grid:      grid,
		dimY:      dimY,
		dimX:      len(lines[0]),
		locations: locations,
		numLocs:   numLocs,
	}
}

func Day24(puzzle Day24Puzzle, part1 bool) uint {
	// Pre-allocate reusable visited array and queue for BFS
	visited := make([][]bool, puzzle.dimY)
	for y := range visited {
		visited[y] = make([]bool, puzzle.dimX)
	}
	queue := make([]image.Point, 0, puzzle.dimX*puzzle.dimY)

	// Calculate distances between all pairs using fixed array
	// distances[from][to] = distance
	var distances [10][10]uint

	for from := 0; from < puzzle.numLocs; from++ {
		bfsDistancesReuse(puzzle, byte(from), visited, queue, &distances)
	}

	// Use DP with bitmasks to solve TSP
	// memo[current][visited_mask] = min distance
	allVisited := uint(1<<puzzle.numLocs) - 1
	memo := make([][1024]uint, 10) // up to 10 locations, 2^10 = 1024 states
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = math.MaxUint
		}
	}

	var tsp func(current int, visited uint) uint
	tsp = func(current int, visited uint) uint {
		if visited == allVisited {
			if part1 {
				return 0
			}
			return distances[current][0]
		}

		if memo[current][visited] != math.MaxUint {
			return memo[current][visited]
		}

		minDist := uint(math.MaxUint)
		for target := 0; target < puzzle.numLocs; target++ {
			bit := uint(1) << target
			if visited&bit == 0 {
				dist := distances[current][target] + tsp(target, visited|bit)
				if dist < minDist {
					minDist = dist
				}
			}
		}

		memo[current][visited] = minDist
		return minDist
	}

	return tsp(0, 1) // start at 0, with 0 visited
}

func bfsDistancesReuse(puzzle Day24Puzzle, start byte, visited [][]bool, queue []image.Point, distances *[10][10]uint) {
	// Clear visited array
	for y := range visited {
		for x := range visited[y] {
			visited[y][x] = false
		}
	}

	startIdx := int(start)
	startPos := puzzle.locations[startIdx]

	queue = queue[:0]
	queue = append(queue, startPos)
	visited[startPos.Y][startPos.X] = true

	dist := uint(0)
	found := 0

	directions := [4]image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for len(queue) > 0 && found < puzzle.numLocs {
		size := len(queue)

		for i := 0; i < size; i++ {
			current := queue[i]

			cell := puzzle.grid[current.Y][current.X]
			if cell >= '0' && cell <= '9' {
				targetIdx := int(cell - '0')
				distances[startIdx][targetIdx] = dist
				found++
			}

			for _, dir := range directions {
				nx, ny := current.X+dir.X, current.Y+dir.Y

				if nx >= 0 && nx < puzzle.dimX &&
					ny >= 0 && ny < puzzle.dimY &&
					!visited[ny][nx] &&
					puzzle.grid[ny][nx] != '#' {
					visited[ny][nx] = true
					queue = append(queue, image.Point{X: nx, Y: ny})
				}
			}
		}

		queue = queue[size:]
		dist++
	}
}
