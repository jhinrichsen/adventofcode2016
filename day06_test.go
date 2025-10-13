package adventofcode2016

import (
	"testing"
)

func TestDay06Part1Example(t *testing.T) {
	const (
		want  = "easter"
		part1 = true
	)
	lines := linesFromFilename(t, exampleFilename(6))
	got := Day06(lines, part1)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay06Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(6))
	for b.Loop() {
		Day06(lines, true)
	}
}

func TestDay06Part2Example(t *testing.T) {
	const (
		want  = "advent"
		part1 = false
	)
	lines := linesFromFilename(t, exampleFilename(6))
	got := Day06(lines, part1)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay06Part1(t *testing.T) {
	const want = "tzstqsua"
	lines := linesFromFilename(t, filename(6))
	got := Day06(lines, true)
	if got != want {
		t.Errorf("Day06() = %v, want %v", got, want)
	}
}

func TestDay06Part2(t *testing.T) {
	const want = "myregdnr"
	lines := linesFromFilename(t, filename(6))
	got := Day06(lines, false)
	if got != want {
		t.Errorf("Day06() = %v, want %v", got, want)
	}
}

func BenchmarkDay06Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(6))
	for b.Loop() {
		Day06(lines, false)
	}
}
