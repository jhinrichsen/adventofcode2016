package adventofcode2016

import (
	"testing"
)

func TestDay03Part1(t *testing.T) {
	const want = 982
	lines := linesFromFilename(t, filename(3))
	got := Day03(lines, true)
	if got != want {
		t.Errorf("Day03() = %v, want %v", got, want)
	}
}

func TestDay03Part2(t *testing.T) {
	const want = 1826
	lines := linesFromFilename(t, filename(3))
	got := Day03(lines, false)
	if got != want {
		t.Errorf("Day03() = %v, want %v", got, want)
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(3))
	for b.Loop() {
		Day03(lines, true)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(3))
	for b.Loop() {
		Day03(lines, false)
	}
}
