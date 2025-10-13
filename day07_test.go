package adventofcode2016

import (
	"testing"
)

func TestDay07Examples(t *testing.T) {
	tests := []struct {
		in    string
		part1 bool
		out   uint
	}{
		{"abba[mnop]qrst", true, 1},
		{"abcd[bddb]xyyx", true, 0},
		{"aaaa[qwer]tyui", true, 0},
		{"ioxxoj[asdfgh]zxcvbn", true, 1},
		{"aba[bab]xyz", false, 1},
		{"xyx[xyx]xyx", false, 0},
		{"aaa[kek]eke", false, 1},
		{"zazbz[bzb]cdb", false, 1},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			want := tt.out
			got := Day07([]string{tt.in}, tt.part1)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(7))
	for b.Loop() {
		Day07(lines, true)
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(7))
	for b.Loop() {
		Day07(lines, false)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = 118
	lines := linesFromFilename(t, filename(7))
	got := Day07(lines, true)
	if got != want {
		t.Errorf("Day07() = %v, want %v", got, want)
	}
}

func TestDay07Part2(t *testing.T) {
	const want = 260
	lines := linesFromFilename(t, filename(7))
	got := Day07(lines, false)
	if got != want {
		t.Errorf("Day07() = %v, want %v", got, want)
	}
}
