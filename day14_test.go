package adventofcode2016

import (
	"testing"
)

func TestDay14Example(t *testing.T) {
	const (
		salt = "abc"
		want = 22728
	)
	got := Day14(salt)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part1(t *testing.T) {
	const (
		salt = "ngcjuoqr"
		want = 18626
	)
	got := Day14(salt)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}

}
