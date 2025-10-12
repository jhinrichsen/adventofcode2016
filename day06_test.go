package adventofcode2016

import (
	"testing"
)

func TestDay6ExamplePart1(t *testing.T) {
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

func TestDay6ExamplePart2(t *testing.T) {
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
