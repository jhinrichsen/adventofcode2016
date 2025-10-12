package adventofcode2016

import (
	"testing"
)

func day23(t *testing.T, part1 bool, filename string, want int) {
	lines := linesFromFilename(t, filename)
	puzzle, err := NewDay23(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(puzzle, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1Example(t *testing.T) {
	day23(t, true, exampleFilename(23), 3)
}

func TestDay23Part1(t *testing.T) {
	day23(t, true, filename(23), 11478)
}

func TestDay23Part2(t *testing.T) {
	day23(t, false, filename(23), 479008038)
}

func BenchmarkDay23Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	puzzle, err := NewDay23(lines)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day23(puzzle, true)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	puzzle, err := NewDay23(lines)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day23(puzzle, false)
	}
}
