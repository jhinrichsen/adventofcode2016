package adventofcode2016

import (
	"testing"
)

func TestDay16Part1Example(t *testing.T) {
	const (
		input    = "10000"
		disksize = 20
		want     = "01100"
	)
	got := Day16(input, disksize)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay16Part1(t *testing.T) {
	const (
		input    = "10011111011011001"
		disksize = 272
		want     = "10111110010110110"
	)
	got := Day16(input, disksize)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay16Part2(t *testing.T) {
	const (
		input    = "10011111011011001"
		disksize = 35651584
		want     = "01101100001100100"
	)
	got := Day16(input, disksize)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}
