package adventofcode2016

import (
	"fmt"
	"strings"
	"testing"
)

var day9Examples = []struct {
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

func TestDay9Examples(t *testing.T) {
	for _, tt := range day9Examples {
		id := fmt.Sprintf("%s (part %d)", tt.in, tt.part)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			var f func(string) uint
			if tt.part == 1 {
				f = Day9Part1
			} else {
				f = Day9Part2
			}
			got := f(tt.in)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay9Part1(b *testing.B) {
	const (
		want  = 150914
		part1 = true
	)
	benchDay9(b, part1, want)
}

func BenchmarkDay9Part2(b *testing.B) {
	const (
		want  = 11052855125
		part1 = false
	)
	benchDay9(b, part1, want)
}

func benchDay9(b *testing.B, part1 bool, want uint) {
	lines, err := linesFromFilename(filename(9))
	if err != nil {
		b.Fatal(err)
	}
	s := strings.Join(lines, "")

	var f func(string) uint
	if part1 {
		f = Day9Part1
	} else {
		f = Day9Part2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := f(s)
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
