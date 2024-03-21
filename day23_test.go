package adventofcode2016

import (
	"testing"
)

func day23(t *testing.T, part1 bool, filename string, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day23(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23ExamplePart1(t *testing.T) {
	day23(t, true, exampleFilename(23), 42)
}

func TestDay23Part1(t *testing.T) {
	day23(t, true, filename(23), 317993)
}

func TestDay23Part2(t *testing.T) {
	day23(t, false, filename(23), 9227647)
}

func BenchmarkDay23Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day23(lines, true)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day23(lines, false)
	}
}
