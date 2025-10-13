package adventofcode2016

import (
	"testing"
)

func TestDay19HighestOneBitValue(t *testing.T) {
	const want = 16
	got := highestOneBitValue(19)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part1Example(t *testing.T) {
	const (
		elves = 5
		want  = 3
	)
	got := Day19Part1(elves)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay19Part1FirstTen(t *testing.T) {
	wants := []uint{
		0,                            // padding
		1, 1, 3, 1, 3, 5, 7, 1, 3, 5, // first entries from OEIS A032434
	}
	for i := 3; i < len(wants); i++ {
		want := wants[i]
		got := Day19Part1(uint(i))
		if want != got {
			t.Fatalf("n=%d want %d but got %d", i+1, want, got)
		}
	}
}

func TestDay19Part1(t *testing.T) {
	const (
		elves = 3001330
		want  = 1808357
	)
	got := Day19Part1(elves)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay19Part1(b *testing.B) {
	for b.Loop() {
		Day19Part1(5)
	}
}

func TestDay19Part2FirstTen(t *testing.T) {
	wants := []uint{
		0,                         // padding
		1, 1, 3, 1, 2, 3, 5, 7, 9, // first entries from OEIS A334473
		1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
	}
	for i := 1; i < len(wants); i++ {
		want := wants[i]
		got := Day19Part2(uint(i))
		if want != got {
			t.Fatalf("n=%d want %d but got %d", i, want, got)
		}
	}
}

func TestDay19Part2(t *testing.T) {
	const want = 1407007
	got := Day19Part2(3001330)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay19Part2(b *testing.B) {
	for b.Loop() {
		Day19Part2(3001330)
	}
}
