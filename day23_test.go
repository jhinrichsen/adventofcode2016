package adventofcode2016

import (
	"testing"
)

func TestDay23Part1Example(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(23))
	puzzle, err := NewDay23(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(puzzle, true)
	if got != 3 {
		t.Fatalf("want %d but got %d", 3, got)
	}
}

func TestDay23Part1(t *testing.T) {
	lines := linesFromFilename(t, filename(23))
	puzzle, err := NewDay23(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(puzzle, true)
	if got != 11478 {
		t.Fatalf("want %d but got %d", 11478, got)
	}
}

func TestDay23Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(23))
	puzzle, err := NewDay23(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(puzzle, false)
	if got != 479008038 {
		t.Fatalf("want %d but got %d", 479008038, got)
	}
}

func BenchmarkDay23Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	puzzle, err := NewDay23(lines)
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		Day23(puzzle, true)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	puzzle, err := NewDay23(lines)
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		Day23(puzzle, false)
	}
}
