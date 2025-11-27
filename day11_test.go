package adventofcode2016

import (
	"testing"
)

func TestDay11Part1(t *testing.T) {
	const (
		want  = 47
		part1 = true
	)
	got, err := Day11(part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part2(t *testing.T) {
	const (
		want  = 71
		part1 = false
	)
	got, err := Day11(part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	for b.Loop() {
		_, _ = Day11(true)
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	for b.Loop() {
		_, _ = Day11(false)
	}
}
