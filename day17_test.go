package adventofcode2016

import (
	"testing"
)

const maxMoves = 1000

var day17Examples = []struct {
	in    string
	part1 bool
	out   string
}{
	{"ihgpwlah", true, "DDRRRD"},
	{"kglvqrro", true, "DDUDRLRRUDRD"},
	{"ulqzkmiv", true, "DRURDRUDDLLDLUURRDULRLDUUDDDRR"},
}

func TestMD5(t *testing.T) {
	// Source: https://en.wikipedia.org/wiki/MD5#MD5_hashes
	const want = "9e107d9d372bb6826bd81d3542a419d6"
	got := md5s("The quick brown fox jumps over the lazy dog")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay17Examples(t *testing.T) {
	for _, tt := range day17Examples {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Day17(tt.in, maxMoves, tt.part1)
			if want != got {
				t.Fatalf("want %q but got %q", want, got)
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
