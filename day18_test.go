package adventofcode2016

import (
	"testing"
)

const (
	input  = "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^"
	repeat = 400000
	want   = 19991126
)

var day18Example = []string{
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
	const want = 1974
	got := Day18(input, 40)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	got := Day18(input, repeat)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
