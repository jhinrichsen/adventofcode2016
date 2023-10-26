package adventofcode2016

import (
	"testing"
)

const (
	day18Input = "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^"

	day18RepeatPart1 = 40
	day18WantPart1   = 1974

	day18RepeatPart2 = 400000
	day18WantPart2   = 19991126
)

var day18ExampleInput = []string{
	".^^.^.^^^^",
	"^^^...^..^",
	"^.^^.^.^^.",
	"..^^...^^^",
	".^^^^.^^.^",
	"^^..^.^^..",
	"^^^^..^^^.",
	"^..^^^^.^^",
	".^^^..^.^^",
	"^^.^^^..^^",
}

func TestDay18Part1(t *testing.T) {
	got := Day18(day18Input, day18RepeatPart1)
	if day18WantPart1 != got {
		t.Fatalf("want %d but got %d", day18WantPart1, got)
	}
}

func TestDay18Part2(t *testing.T) {
	got := Day18(day18Input, day18RepeatPart2)
	if day18WantPart2 != got {
		t.Fatalf("want %d but got %d", day18WantPart2, got)
	}
}
