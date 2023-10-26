package adventofcode2016

import (
	"testing"
)

func TestTrapsAsString(t *testing.T) {
	want := "^^.^.^."
	got := trapsAsString([]bool{
		false, // left safe
		true, true, false, true, false, true, false,
		false, // left safe
	})
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay18ExampleV3(t *testing.T) {
	const want = 38
	got := Day18V3(day18ExampleInput[0], len(day18ExampleInput))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part1V3(t *testing.T) {
	got := Day18V3(day18Input, day18RepeatPart1)
	if day18WantPart1 != got {
		t.Fatalf("want %d but got %d", day18WantPart1, got)
	}
}

func BenchmarkDay18Part1V3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day18V3(day18Input, day18RepeatPart1)
	}
}

func TestDay18Part2V3(t *testing.T) {
	got := Day18V3(day18Input, day18RepeatPart2)
	if day18WantPart2 != got {
		t.Fatalf("want %d but got %d", day18WantPart2, got)
	}
}

func BenchmarkDay18Part2V3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day18V3(day18Input, day18RepeatPart2)
	}
}
