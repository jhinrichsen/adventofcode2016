package adventofcode2016

import (
	"testing"
)

func day12(t *testing.T, part1 bool, filename string, want int) {
	lines := linesFromFilename(t, filename)
	got, err := Day12(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part1Example(t *testing.T) {
	day12(t, true, exampleFilename(12), 42)
}

func TestDay12Part1(t *testing.T) {
	day12(t, true, filename(12), 317993)
}

func TestDay12Part2(t *testing.T) {
	day12(t, false, filename(12), 9227647)
}

func BenchmarkDay12Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(12))
	for b.Loop() {
		Day12(lines, true)
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(12))
	for b.Loop() {
		Day12(lines, false)
	}
}
