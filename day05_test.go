package adventofcode2016

import (
	"testing"
)

const day5Input = "uqwqemis"

func TestDay5Part1Example(t *testing.T) {
	const (
		want  = "18f47a30"
		part1 = true
	)

	got := Day5("abc", part1)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay05Part1(t *testing.T) {
	const want = "1a3099aa"
	got := Day5(day5Input, true)
	if got != want {
		t.Errorf("Day5() = %v, want %v", got, want)
	}
}

func TestDay05Part2(t *testing.T) {
	const want = "694190cd"
	got := Day5(day5Input, false)
	if got != want {
		t.Errorf("Day5() = %v, want %v", got, want)
	}
}

func BenchmarkDay5Part1(b *testing.B) {
	const (
		want  = "1a3099aa"
		part1 = true
	)
	for i := 0; i < b.N; i++ {
		got := Day5(day5Input, part1)
		if want != got {
			b.Fatalf("want %q but got %q", want, got)
		}
	}
}

func TestDay5Part2Example(t *testing.T) {
	const (
		want  = "05ace8e3"
		part1 = false
	)
	got := Day5("abc", part1)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay5Part2(b *testing.B) {
	const (
		want  = "694190cd"
		part1 = false
	)
	for i := 0; i < b.N; i++ {
		got := Day5(day5Input, part1)
		if want != got {
			b.Fatalf("want %q but got %q", want, got)
		}
	}
}
