package adventofcode2016

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay09Examples(t *testing.T) {
	tests := []struct {
		in   string
		out  uint
		part uint8
	}{
		// part 1
		{"ADVENT", 6, 1},
		{"A(1x5)BC", 7, 1},
		{"(3x3)XYZ", 9, 1},
		{"A(2x2)BCD(2x2)EFG", 11, 1},
		{"(6x1)(1x3)A", 6, 1},
		{"X(8x2)(3x3)ABCY", 18, 1},

		// part 2
		{"(3x3)XYZ", 9, 2},
		{"X(8x2)(3x3)ABCY", 20, 2},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920, 2},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445, 2},
	}

	for _, tt := range tests {
		id := fmt.Sprintf("%s (part %d)", tt.in, tt.part)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			var f func(string) uint
			if tt.part == 1 {
				f = Day09Part1
			} else {
				f = Day09Part2
			}
			got := f(tt.in)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay09Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(9))
	s := strings.Join(lines, "")
	for b.Loop() {
		Day09Part1(s)
	}
}

func BenchmarkDay09Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(9))
	s := strings.Join(lines, "")
	for b.Loop() {
		Day09Part2(s)
	}
}

func TestDay09Part1(t *testing.T) {
	const want = 150914
	lines := linesFromFilename(t, filename(9))
	s := strings.Join(lines, "")
	got := Day09Part1(s)
	if got != want {
		t.Errorf("Day09Part1() = %v, want %v", got, want)
	}
}

func TestDay09Part2(t *testing.T) {
	const want = 11052855125
	lines := linesFromFilename(t, filename(9))
	s := strings.Join(lines, "")
	got := Day09Part2(s)
	if got != want {
		t.Errorf("Day09Part2() = %v, want %v", got, want)
	}
}

