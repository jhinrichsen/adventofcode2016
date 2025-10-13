package adventofcode2016

import (
	"testing"
)

func TestDay05Part1Example(t *testing.T) {
	const want = "18f47a30"
	got := Day05("abc", true)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay05Part1(t *testing.T) {
	const want = "1a3099aa"
	got := Day05("uqwqemis", true)
	if got != want {
		t.Errorf("Day05() = %v, want %v", got, want)
	}
}

func TestDay05Part2(t *testing.T) {
	const want = "694190cd"
	got := Day05("uqwqemis", false)
	if got != want {
		t.Errorf("Day05() = %v, want %v", got, want)
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	for b.Loop() {
		Day05("uqwqemis", true)
	}
}

func TestDay05Part2Example(t *testing.T) {
	const want = "05ace8e3"
	got := Day05("abc", false)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	for b.Loop() {
		Day05("uqwqemis", false)
	}
}
