package adventofcode2016

import (
	"testing"
)

func day12v2(t *testing.T, part1 bool, filename string, want int) {
	lines := linesFromFilename(t, filename)
	got, err := Day12V2(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12V2Part1Example(t *testing.T) {
	day12v2(t, true, exampleFilename(12), 42)
}

func TestDay12V2Part1(t *testing.T) {
	day12v2(t, true, filename(12), 317993)
}

func TestDay12V2Part2(t *testing.T) {
	day12v2(t, false, filename(12), 9227647)
}

func BenchmarkDay12V2Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(12))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day12V2(lines, true)
	}
}

func BenchmarkDay12V2Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(12))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day12V2(lines, false)
	}
}
