package adventofcode2016

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDay22Part1(t *testing.T) {
	const (
		part1 = true
		// want  = 20636 // first try wrong, too high
		// i was using col 1 (size) and col 2 (used) instead of col 2 (used) and col 3 (avail)
		// does that count as 'off-by-one'? not sure...
		want = 937
	)
	lines := linesFromFilename(t, filename(22))
	got, err := Day22(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part2(t *testing.T) {
	const (
		part1 = false
		// want  = 181 // first try wrong, too low
		// want = 187 // fixing dim+1, too low
		want = 188 // counting moves instead of images
	)
	lines := linesFromFilename(t, filename(22))
	got, err := Day22(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay22Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(22))
	for b.Loop() {
		Day22(lines, true)
	}
}

func BenchmarkDay22Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(22))
	for b.Loop() {
		Day22(lines, false)
	}
}

func TestDay22Part2Visualization(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping visualization test in short mode")
	}

	const outputPath = "img/day22.gif"

	// Skip if file already exists
	if _, err := os.Stat(outputPath); err == nil {
		t.Skipf("visualization already exists at %s (delete to regenerate)", outputPath)
	}

	lines := linesFromFilename(t, filename(22))

	// Parse grid (copied from Day22 function)
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

	parseHuman := func(s string) (uint64, error) {
		idx := len(s) - 1
		n, err := strconv.Atoi(s[:idx])
		if err != nil {
			return 0, err
		}
		var unit uint64
		const (
			KB = 1024
			MB = KB * KB
			GB = MB * KB
			TB = GB * KB
			PB = TB * KB
		)

		switch s[idx:][0] {
		case 'K':
			unit = KB
		case 'M':
			unit = MB
		case 'T':
			unit = TB
		default:
			return 0, nil
		}
		return uint64(n) * unit, nil
	}

	type coordinate struct {
		x, y int
	}
	type df struct {
		used, avail uint64
	}

	m := make(map[coordinate]df)
	var dim, empty coordinate

	for _, line := range lines {
		if line[0] != '/' {
			continue
		}
		parts := strings.Fields(line)
		x, y, err := parseNode(strings.Split(parts[0], "/")[3])
		if err != nil {
			t.Fatal(err)
		}

		dim.x = max(dim.x, x)
		dim.y = max(dim.y, y)

		used, err := parseHuman(parts[2])
		if err != nil {
			t.Fatal(err)
		}

		avail, err := parseHuman(parts[3])
		if err != nil {
			t.Fatal(err)
		}
		m[coordinate{x, y}] = df{used, avail}
	}

	dim.x++
	dim.y++

	// Find empty node
	for ca, na := range m {
		if na.used == 0 {
			empty.x = ca.x
			empty.y = ca.y
			break
		}
	}

	// normalize used% to 255 grayscales
	per8 := make(map[coordinate]uint8, dim.x*dim.y)
	for k, v := range m {
		n := uint8((255 * v.used) / (v.used + v.avail))
		per8[k] = n
	}
	// top right pixel is our special index 1
	per8[coordinate{dim.x - 1, 0}] = 1

	// convert gray to RGBA
	var palette color.Palette
	for i := 0; i < 256; i++ {
		palette = append(palette, color.Gray{Y: uint8(i)})
	}
	// index 1 is red
	palette[1] = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	// create image
	type path struct {
		direction coordinate
		n         int
	}
	var (
		images []*image.Paletted
		delays []int

		rect  = image.Rect(0, 0, dim.x, dim.y)
		left  = coordinate{-1, 0}
		right = coordinate{+1, 0}
		up    = coordinate{0, -1}
		down  = coordinate{0, +1}
		paths = []path{
			{left, 4},
			{up, 22},
			{right, 22},
		}
		next coordinate
	)

	// repeat sequence to move red hole to the left
	for i := 0; i < dim.x-2; i++ {
		paths = append(paths, []path{
			{down, 1},
			{left, 2},
			{up, 1},
			{right, 1},
		}...)
	}
	for i := 0; i < len(paths); i++ {
		direction := paths[i].direction
		n := paths[i].n
		for j := 0; j < n; j++ {
			next.x = empty.x + direction.x
			next.y = empty.y + direction.y

			// swap empty and next
			tmp := per8[next]
			per8[next] = per8[empty]
			per8[empty] = tmp
			empty = next

			g := image.NewPaletted(rect, palette)
			for k, v := range per8 {
				g.SetColorIndex(k.x, k.y, v)
			}
			images = append(images, g)
			delays = append(delays, 5) // 5/100 = 0.05 seconds = 50ms per frame
		}
	}
	// Create img directory if it doesn't exist
	if err := os.MkdirAll("img", 0755); err != nil {
		t.Fatal(err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	err = gif.EncodeAll(f, &gif.GIF{Image: images, Delay: delays})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Generated visualization at %s", outputPath)
}

func TestDay22Part2VisualizationAPNG(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping visualization test in short mode")
	}

	const outputPath = "img/day22.apng"

	// Skip if file already exists
	if _, err := os.Stat(outputPath); err == nil {
		t.Skipf("visualization already exists at %s (delete to regenerate)", outputPath)
	}

	lines := linesFromFilename(t, filename(22))

	// Parse grid (copied from Day22 function)
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

	parseHuman := func(s string) (uint64, error) {
		idx := len(s) - 1
		n, err := strconv.Atoi(s[:idx])
		if err != nil {
			return 0, err
		}
		var unit uint64
		const (
			KB = 1024
			MB = KB * KB
			GB = MB * KB
			TB = GB * KB
			PB = TB * KB
		)

		switch s[idx:][0] {
		case 'K':
			unit = KB
		case 'M':
			unit = MB
		case 'T':
			unit = TB
		default:
			return 0, nil
		}
		return uint64(n) * unit, nil
	}

	type coordinate struct {
		x, y int
	}
	type df struct {
		used, avail uint64
	}

	m := make(map[coordinate]df)
	var dim, empty coordinate

	for _, line := range lines {
		if line[0] != '/' {
			continue
		}
		parts := strings.Fields(line)
		x, y, err := parseNode(strings.Split(parts[0], "/")[3])
		if err != nil {
			t.Fatal(err)
		}

		dim.x = max(dim.x, x)
		dim.y = max(dim.y, y)

		used, err := parseHuman(parts[2])
		if err != nil {
			t.Fatal(err)
		}

		avail, err := parseHuman(parts[3])
		if err != nil {
			t.Fatal(err)
		}
		m[coordinate{x, y}] = df{used, avail}
	}

	dim.x++
	dim.y++

	// Find empty node
	for ca, na := range m {
		if na.used == 0 {
			empty.x = ca.x
			empty.y = ca.y
			break
		}
	}

	// normalize used% to 255 grayscales
	per8 := make(map[coordinate]uint8, dim.x*dim.y)
	for k, v := range m {
		n := uint8((255 * v.used) / (v.used + v.avail))
		per8[k] = n
	}
	// top right pixel is our special index 1
	per8[coordinate{dim.x - 1, 0}] = 1

	// create image
	type path struct {
		direction coordinate
		n         int
	}
	var (
		rect  = image.Rect(0, 0, dim.x, dim.y)
		left  = coordinate{-1, 0}
		right = coordinate{+1, 0}
		up    = coordinate{0, -1}
		down  = coordinate{0, +1}
		paths = []path{
			{left, 4},
			{up, 22},
			{right, 22},
		}
		next coordinate
	)

	// repeat sequence to move red hole to the left
	for i := 0; i < dim.x-2; i++ {
		paths = append(paths, []path{
			{down, 1},
			{left, 2},
			{up, 1},
			{right, 1},
		}...)
	}

	// Create img directory if it doesn't exist
	if err := os.MkdirAll("img", 0755); err != nil {
		t.Fatal(err)
	}

	// Generate frames as a sequence
	seq := func(yield func(image.Image) bool) {
		for i := 0; i < len(paths); i++ {
			direction := paths[i].direction
			n := paths[i].n
			for j := 0; j < n; j++ {
				next.x = empty.x + direction.x
				next.y = empty.y + direction.y

				// swap empty and next
				tmp := per8[next]
				per8[next] = per8[empty]
				per8[empty] = tmp
				empty = next

				// Create RGBA image for better quality
				img := image.NewRGBA(rect)
				for k, v := range per8 {
					var c color.RGBA
					if v == 1 {
						// Red for goal data
						c = color.RGBA{R: 255, G: 0, B: 0, A: 255}
					} else {
						// Grayscale for usage
						c = color.RGBA{R: v, G: v, B: v, A: 255}
					}
					img.Set(k.x, k.y, c)
				}
				if !yield(img) {
					return
				}
			}
		}
	}

	// Constant 50ms delay (num=5, den=100 means 5/100 = 0.05 seconds)
	delay := func(i int) (num, den uint16) {
		return 5, 100
	}

	// Encode as APNG with infinite loop
	if err := PackAPNGFileFromSeq(seq, math.MaxUint64, outputPath, delay); err != nil {
		t.Fatal(err)
	}

	t.Logf("Generated APNG visualization at %s", outputPath)
}
