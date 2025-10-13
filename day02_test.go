package adventofcode2016

import (
	"testing"
)

func TestDay02Part1Example(t *testing.T) {
	tests := []struct {
		part1 bool
		want  string
	}{
		{true, "1985"},
		{false, "5DB3"},
	}

	for _, tt := range tests {
		input := exampleFile(t, 2)
		got := Day02(input, tt.part1)
		if tt.want != got {
			t.Fatalf("want %q but got %q", tt.want, got)
		}
	}
}

func TestDay02Part1(t *testing.T) {
	const want = "65556"
	input := file(t, 2)
	got := Day02(input, true)
	if got != want {
		t.Errorf("Day02() = %v, want %v", got, want)
	}
}

func TestDay02Part2(t *testing.T) {
	const want = "CB779"
	input := file(t, 2)
	got := Day02(input, false)
	if got != want {
		t.Errorf("Day02() = %v, want %v", got, want)
	}
}

func BenchmarkDay02Part1(b *testing.B) {
	input := file(b, 2)
	for b.Loop() {
		Day02(input, true)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	input := file(b, 2)
	for b.Loop() {
		Day02(input, false)
	}
}

