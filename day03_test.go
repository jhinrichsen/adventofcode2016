package adventofcode2016

import (
	"testing"
)

func BenchmarkDay3Part1(b *testing.B) {
	const want = 982
	benchDay3(b, true, want)
}

func BenchmarkDay3Part2(b *testing.B) {
	const want = 1826
	benchDay3(b, false, want)
}

func benchDay3(b *testing.B, part1 bool, want uint) {
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := Day3(lines, part1)
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
