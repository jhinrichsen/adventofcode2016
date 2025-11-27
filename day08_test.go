package adventofcode2016

import (
	"strings"
	"testing"

	"gitlab.com/jhinrichsen/aococr"
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
	screen := newDay08(width, height)
	got, err := Day08(screen, lines, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Day08() = %v, want %v", got, want)
	}
}

func TestDay08Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(8))
	screen := newDay08(width, height)
	_, err := Day08(screen, lines, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	display := screen.String()

	charSet := map[rune]bool{'#': true, '.': false}
	got, err := aococr.ParseLetters(display, charSet)
	if err != nil {
		t.Fatal(err)
	}

	const want = "EOARGPHYAO"
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay08Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(8))
	for b.Loop() {
		_, _ = Day08(newDay08(width, height), lines, true, nil)
	}
}

func BenchmarkDay08Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(8))
	charSet := map[rune]bool{'#': true, '.': false}
	for b.Loop() {
		screen := newDay08(width, height)
		_, _ = Day08(screen, lines, false, nil)
		_, _ = aococr.ParseLetters(screen.String(), charSet)
	}
}
