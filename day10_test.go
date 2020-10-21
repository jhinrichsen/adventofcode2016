package adventofcode2016

import (
	"testing"
)

// The example provided allows to find the right bot by just looking at the
// syntax. Part #1 does not allow the same, one needs to really analyse the
// complete chain.
func TestDay10Example(t *testing.T) {
	const want = 2
	lines, err := linesFromFilename(exampleFilename(10))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day10(lines, 2, 5)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	const want = 161
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day10(lines, 61, 17)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
