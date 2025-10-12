package adventofcode2016

import (
	"testing"
)

var day7Examples = []struct {
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

func TestDay7Examples(t *testing.T) {
	for _, tt := range day7Examples {
		t.Run(tt.in, func(t *testing.T) {
			want := tt.out
			got := Day7([]string{tt.in}, tt.part1)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay7Part1(b *testing.B) {
	const (
		want  = 118
		part1 = true
	)
	benchDay7(b, part1, want)
}

func BenchmarkDay7Part2(b *testing.B) {
	const (
		want  = 260
		part1 = false
	)
	benchDay7(b, part1, want)
}

func TestDay7Part1(t *testing.T) {
	const want = 118
	lines := linesFromFilename(t, filename(7))
	got := Day7(lines, true)
	if got != want {
		t.Errorf("Day7() = %v, want %v", got, want)
	}
}

func TestDay7Part2(t *testing.T) {
	const want = 260
	lines := linesFromFilename(t, filename(7))
	got := Day7(lines, false)
	if got != want {
		t.Errorf("Day7() = %v, want %v", got, want)
	}
}

func benchDay7(b *testing.B, part1 bool, want uint) {
	lines := linesFromFilename(b, filename(7))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := Day7(lines, part1)
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
