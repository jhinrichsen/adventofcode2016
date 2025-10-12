package adventofcode2016

import (
	"testing"
)

func TestDay6Part1Example(t *testing.T) {
	const (
		want  = "easter"
		part1 = true
	)
	lines := linesFromFilename(t, exampleFilename(6))
	got := Day6(lines, part1)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay6Part1(b *testing.B) {
	const (
		want  = "tzstqsua"
		part1 = true
	)
	lines := linesFromFilename(b, filename(6))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := Day6(lines, part1)
		if want != got {
			b.Fatalf("want %q but got %q", want, got)
		}
	}
}

func TestDay6Part2Example(t *testing.T) {
	const (
		want  = "advent"
		part1 = false
	)
	lines := linesFromFilename(t, exampleFilename(6))
	got := Day6(lines, part1)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay6Part1(t *testing.T) {
	const want = "tzstqsua"
	lines := linesFromFilename(t, filename(6))
	got := Day6(lines, true)
	if got != want {
		t.Errorf("Day6() = %v, want %v", got, want)
	}
}

func TestDay6Part2(t *testing.T) {
	const want = "myregdnr"
	lines := linesFromFilename(t, filename(6))
	got := Day6(lines, false)
	if got != want {
		t.Errorf("Day6() = %v, want %v", got, want)
	}
}

func BenchmarkDay6Part2(b *testing.B) {
	const (
		want  = "myregdnr"
		part1 = false
	)
	lines := linesFromFilename(b, filename(6))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := Day6(lines, part1)
		if want != got {
			b.Fatalf("want %q but got %q", want, got)
		}
	}
}
