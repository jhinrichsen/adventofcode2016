package adventofcode2016

import (
	"testing"
)

func TestDay14Example(t *testing.T) {
	const (
		salt  = "abc"
		part1 = true
		want  = 22728
	)
	got := Day14(salt, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part1(t *testing.T) {
	const (
		salt  = "ngcjuoqr"
		part1 = true
		want  = 18626
	)
	got := Day14(salt, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay14Part1(b *testing.B) {
	for b.Loop() {
		Day14("ngcjuoqr", true)
	}
}

func TestStretchedHash(t *testing.T) {
	const (
		salt = "abc"
		want = "a107ff634856bb300138cac6568c0f24"
	)
	got := stretchedHash(salt + "0")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay14Part2(t *testing.T) {
	const (
		salt  = "ngcjuoqr"
		part1 = false
		want  = 20092
	)
	got := Day14(salt, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
