package adventofcode2016

import (
	"testing"
)

const maxMoves = 1000

var day17Examples = []struct {
	in       string
	shortest string
	longest  int
}{
	{"ihgpwlah", "DDRRRD", 370},
	{"kglvqrro", "DDUDRLRRUDRD", 492},
	{"ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR", 830},
}

func TestMD5(t *testing.T) {
	// Source: https://en.wikipedia.org/wiki/MD5#MD5_hashes
	const want = "9e107d9d372bb6826bd81d3542a419d6"
	got := md5s("The quick brown fox jumps over the lazy dog")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay17ExamplesPart1(t *testing.T) {
	for _, tt := range day17Examples {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.shortest
			got := Day17(tt.in, maxMoves, true)
			if want != got {
				t.Fatalf("want %q but got %q", want, got)
			}
		})
	}
}

func TestDay17ExamplesPart2(t *testing.T) {
	for _, tt := range day17Examples {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.longest
			got := len(Day17(tt.in, tt.longest+2, false))
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay17Part1(t *testing.T) {
	const (
		input = "edjrjqaa"
		part1 = true
		want  = "DUDRDLRRRD"
	)
	got := Day17(input, maxMoves, part1)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay17Part2(t *testing.T) {
	const (
		want  = 502
		input = "edjrjqaa"
		part1 = false
	)
	got := len(Day17(input, maxMoves, part1))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
