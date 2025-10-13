package adventofcode2016

import (
	"testing"
)

func TestDay25(t *testing.T) {
	const want = 158
	puzzle, err := NewDay25(linesFromFilename(t, filename(25)))
	if err != nil {
		t.Fatal(err)
	}
	got := Day25(puzzle)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay25(b *testing.B) {
	puzzle, err := NewDay25(linesFromFilename(b, filename(25)))
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		Day25(puzzle)
	}
}
