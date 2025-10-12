package adventofcode2016

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

var day8Steps = [][]string{
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

func TestDay8Example(t *testing.T) {
	const (
		want  = 6
		part1 = true
	)
	lines := linesFromFilename(t, exampleFilename(8))
	screen := newDay8(7, 3)
	var i uint
	f := func(d day8) {
		want := strings.Join(day8Steps[i], "\n")
		got := d.String()
		if want != got {
			t.Fatalf("want %q but got %q", want, got)
		}
		i++
	}
	got, err := Day8(screen, lines, part1, f)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay8Part1(t *testing.T) {
	const want = 128
	lines := linesFromFilename(t, filename(8))
	got, err := Day8(newDay8(width, height), lines, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Day8() = %v, want %v", got, want)
	}
}

func BenchmarkDay8Part1(b *testing.B) {
	const (
		want  = 128
		part1 = true
	)
	lines := linesFromFilename(b, filename(8))
	b.ResetTimer()
	// save last screen
	var save day8
	for i := 0; i < b.N; i++ {
		got, err := Day8(newDay8(width, height), lines, part1, func(d day8) {
			save = d
		})
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
	/*
	   ####   ##    ##   ###    ##   ###   #  #  #   #  ##    ##
	   #     #  #  #  #  #  #  #  #  #  #  #  #  #   # #  #  #  #
	   ###   #  #  #  #  #  #  #     #  #  ####   # #  #  #  #  #
	   #     #  #  ####  ###   # ##  ###   #  #    #   ####  #  #
	   #     #  #  #  #  # #   #  #  #     #  #    #   #  #  #  #
	   ####   ##   #  #  #  #   ###  #     #  #    #   #  #   ##
	*/
	fmt.Fprint(io.Discard, strings.ReplaceAll(save.String(), ".", " "))
}
