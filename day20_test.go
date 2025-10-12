package adventofcode2016

import (
	"testing"
)

func TestDay20Part1Example(t *testing.T) {
	const (
		day   = 20
		part1 = true
		want  = 3
	)
	lines := linesFromFilename(t, exampleFilename(day))
	got, err := Day20(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part2Example(t *testing.T) {
	const (
		day   = 20
		part1 = false
		want  = 1
	)
	lines := linesFromFilename(t, exampleFilename(day))
	got, err := Day20(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part1(t *testing.T) {
	const (
		day   = 20
		part1 = true
		want  = 23923783
	)
	lines := linesFromFilename(t, filename(day))
	got, err := Day20(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Part1(b *testing.B) {
	const (
		day   = 20
		part1 = true
	)
	lines := linesFromFilename(b, filename(day))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day20(lines, part1)
	}
}

func TestDay20Part2(t *testing.T) {
	const (
		day   = 20
		part1 = false
		want  = 3195615
	)
	lines := linesFromFilename(t, filename(day))
	got, err := Day20(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Part2(b *testing.B) {
	const (
		day   = 20
		part1 = false
	)
	lines := linesFromFilename(b, filename(day))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day20(lines, part1)
	}
}
