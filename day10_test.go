package adventofcode2016

import (
	"testing"
)

// The example provided allows to find the right bot by just looking at the
// syntax. Part #1 does not allow the same, one needs to really analyse the
// complete chain.
func TestDay10Example(t *testing.T) {
	const (
		want  = 2
		part1 = true
	)
	lines := linesFromFilename(t, exampleFilename(10))
	got, err := Day10(lines, part1, 2, 5)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	const (
		want  = 161
		part1 = true
	)
	lines := linesFromFilename(b, filename(10))
	for i := 0; i < b.N; i++ {
		got, err := Day10(lines, part1, 61, 17)
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	const (
		want  = 133163
		part1 = false
	)
	lines := linesFromFilename(b, filename(10))
	for i := 0; i < b.N; i++ {
		got, err := Day10(lines, part1, 0, 0)
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
