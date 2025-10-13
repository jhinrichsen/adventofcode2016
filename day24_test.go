package adventofcode2016

import (
	"testing"
)

func TestDay24Part1Example(t *testing.T) {
	const want = 14
	got := Day24(NewDay24(linesFromFilename(t, exampleFilename(24))), true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part1(t *testing.T) {
	const want = 490
	got := Day24(NewDay24(linesFromFilename(t, filename(24))), true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part2(t *testing.T) {
	const want = 744
	got := Day24(NewDay24(linesFromFilename(t, filename(24))), false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay24Part1(b *testing.B) {
	puzzle := NewDay24(linesFromFilename(b, filename(24)))
	for b.Loop() {
		Day24(puzzle, true)
	}
}

func BenchmarkDay24Part2(b *testing.B) {
	puzzle := NewDay24(linesFromFilename(b, filename(24)))
	for b.Loop() {
		Day24(puzzle, false)
	}
}
