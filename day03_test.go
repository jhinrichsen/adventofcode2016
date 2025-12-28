package adventofcode2016

import (
	"testing"
)

func TestDay03Part1(t *testing.T) {
	const want = 982
	input := file(t, 3)
	got := Day03(input, true)
	if got != want {
		t.Errorf("Day03() = %v, want %v", got, want)
	}
}

func TestDay03Part2(t *testing.T) {
	const want = 1826
	input := file(t, 3)
	got := Day03(input, false)
	if got != want {
		t.Errorf("Day03() = %v, want %v", got, want)
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	input := file(b, 3)
	for b.Loop() {
		Day03(input, true)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	input := file(b, 3)
	for b.Loop() {
		Day03(input, false)
	}
}
