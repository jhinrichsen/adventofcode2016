package adventofcode2016

import (
	"strings"
	"testing"
)

func TestDay08Example(t *testing.T) {
	steps := [][]string{
		{
			"###....",
			"###....",
			".......",
			"",
		}, {
			"#.#....",
			"###....",
			".#.....",
			"",
		}, {
			"....#.#",
			"###....",
			".#.....",
			"",
		}, {
			".#..#.#",
			"#.#....",
			".#.....",
			"",
		},
	}
	const (
		want  = 6
		part1 = true
	)
	lines := linesFromFilename(t, exampleFilename(8))
	screen := newDay08(7, 3)
	var i uint
	f := func(d day8) {
		want := strings.Join(steps[i], "\n")
		got := d.String()
		if want != got {
			t.Fatalf("want %q but got %q", want, got)
		}
		i++
	}
	got, err := Day08(screen, lines, part1, f)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part1(t *testing.T) {
	const want = 128
	lines := linesFromFilename(t, filename(8))
	got, err := Day08(newDay08(width, height), lines, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Day08() = %v, want %v", got, want)
	}
}

func BenchmarkDay08Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(8))
	for b.Loop() {
		Day08(newDay08(width, height), lines, true, nil)
	}
}
