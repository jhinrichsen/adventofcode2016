package adventofcode2016

import (
	"testing"
)

func TestDay18Example(t *testing.T) {
	const (
		in   = "..^^."
		want = ".^^^^"
	)
	got := next(in)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay18ExampleV1(t *testing.T) {
	const want = 38
	for i := 1; i < len(day18ExampleInput); i++ {
		want := day18ExampleInput[i]
		got := next(day18ExampleInput[i-1])
		if want != got {
			t.Fatalf("want %q but got %q", want, got)
		}
	}

	got := Day18V1(day18ExampleInput[0], len(day18ExampleInput))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2V1(t *testing.T) {
	got := Day18(day18Input, day18RepeatPart2)
	if day18WantPart2 != got {
		t.Fatalf("want %d but got %d", day18WantPart2, got)
	}
}

func BenchmarkDay18Part2V1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day18V1(day18Input, day18RepeatPart2)
	}
}
