package adventofcode2016

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
	"strconv"
	"strings"
)

// Day22: Grid Computing
func Day22(lines []string, part1 bool) (uint, error) {
	// 'node-x1-y2' -> 1, 2, nil
	parseNode := func(s string) (int, int, error) {
		parts := strings.Split(s, "-")
		x, err := strconv.Atoi(parts[1][1:])
		if err != nil {
			return 0, 0, err
		}
		y, err := strconv.Atoi(parts[2][1:])
		if err != nil {
			return x, 0, err
		}
		return x, y, nil
	}

	// from 'man df':
	//        -h, --human-readable
	//              print sizes in powers of 1024 (e.g., 1023M)
	//
	//       -H, --si
	//              print sizes in powers of 1000 (e.g., 1.1G)
	// A 64 bit unsigned integer can hole 20 digits, PB has 15 digits.
	parseHuman := func(s string) (uint64, error) {
		idx := len(s) - 1
		n, err := strconv.Atoi(s[:idx])
		if err != nil {
			return 0, err
		}
		var unit uint64
		// our puzzle input uses 'df -h', so:
		const (
			KB = 1024    // 10^3
			MB = KB * KB // 10^6
			GB = MB * KB // 10^9
			TB = GB * KB // 10^12
			PB = TB * KB // 10^15
		)

		switch s[idx:][0] {
		case 'K':
			unit = KB
		case 'M':
			unit = MB
		case 'T':
			unit = TB
		default:
			return 0, fmt.Errorf("unknown unit in %q", s)
		}
		return uint64(n) * unit, nil
	}

	// use a map for now, because dimensions for an array are unknown yet, and see how fast this is.
	type coordinate struct {
		x, y int
	}
	type df struct {
		used, avail uint64
	}
	estimatedNodes := int(math.Sqrt(float64(len(lines)))) // assume square, we could look it up, but that's sort of cheating...
	m := make(map[coordinate]df, estimatedNodes)

	var dim, empty coordinate
	// root@ebhq-gridcenter# df -h
	// Filesystem              Size  Used  Avail  Use%
	// /dev/grid/node-x0-y0     92T   72T    20T   78%
	for i, line := range lines {
		if line[0] != '/' {
			continue
		}
		parts := strings.Fields(line)
		x, y, err := parseNode(strings.Split(parts[0], "/")[3])
		if err != nil {
			return 0, fmt.Errorf("line %d: %w", i+1, err)
		}

		dim.x = max(dim.x, x)
		dim.y = max(dim.y, y)

		used, err := parseHuman(parts[2])
		if err != nil {
			return 0, fmt.Errorf("error parsing line %d: %w", i, err)
		}

		avail, err := parseHuman(parts[3])
		if err != nil {
			return 0, fmt.Errorf("error parsing line %d: %w", i, err)
		}
		m[coordinate{x, y}] = df{used, avail}
	}

	// part 1
	var viable uint
	for ca, na := range m {
		for cb, nb := range m {
			// 'Nodes A and B are not the same node.'
			if ca == cb {
				continue
			}
			// 'Node A is not empty (its Used is not zero).'
			if na.used == 0 {
				empty.x = ca.x
				empty.y = ca.y
				continue
			}
			// 'The data on node A (its Used) would fit on node B (its Avail).'
			if na.used > nb.avail {
				continue
			}
			viable++
		}
	}
	if part1 {
		return viable, nil
	}

	// normalize used% to 8 bit value
	per8 := make(map[coordinate]uint8, dim.x*dim.y)
	for k, v := range m {
		n := (256 * v.used) / (v.used + v.avail)
		per8[k] = uint8(n)
	}

	// convert gray to RGBA
	var palette color.Palette
	for i := 0; i < 256; i++ {
		palette = append(palette, color.Gray{Y: uint8(i)})
	}

	// create image
	var (
		images []*image.Paletted
		delays []int

		rect  = image.Rect(0, 0, dim.x, dim.y)
		left  = coordinate{-1, 0}
		right = coordinate{+1, 0}
		up    = coordinate{0, -1}
		down  = coordinate{0, +1}
		paths = []struct {
			direction coordinate
			n         int
		}{
			{left, 4},
			{up, 22},
			{right, 21},
			{down, 1},
		}
		count uint
		next  coordinate
	)
	for i := 0; i < len(paths); i++ {
		direction := paths[i].direction
		n := paths[i].n
		for j := 0; j < n; j++ {
			next.x = empty.x + direction.x
			next.y = empty.y + direction.y
			tmp := per8[next]
			per8[next] = per8[empty]
			per8[empty] = tmp
			empty = next

			g := image.NewPaletted(rect, palette)
			for k, v := range per8 {
				// TODO SetColorIndex
				g.Set(k.x, k.y, palette[v])

			}
			images = append(images, g)
			delays = append(delays, 100) // 10 * 100ms

			count++
		}
	}
	f, err := os.Create("day22.gif")
	if err != nil {
		return viable, err
	}
	defer f.Close()
	err = gif.EncodeAll(f, &gif.GIF{Image: images, Delay: delays})
	return count, err
}
