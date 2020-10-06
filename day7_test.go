package adventofcode2016

import (
	"fmt"
	"testing"
)

var day7Examples = []struct {
	in  string
	out uint
}{
	{"abba[mnop]qrst", 1},
	{"abcd[bddb]xyyx", 0},
	{"aaaa[qwer]tyui", 0},
	{"ioxxoj[asdfgh]zxcvbn", 1},
}

func TestDay7ExamplesPart1(t *testing.T) {
	for _, tt := range day7Examples {
		id := fmt.Sprintf("%s", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Day7([]string{tt.in})
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay7Part1(b *testing.B) {
	const want = 118
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := Day7(lines)
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
