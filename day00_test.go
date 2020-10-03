package adventofcode2016

import (
	"fmt"
	"testing"
)

var day00Examples = []struct {
	in  []string
	out uint
}{
	{[]string{}, 0},
}

func TestDay00ExamplesPart1(t *testing.T) {
	for _, tt := range day00Examples {
		id := fmt.Sprintf("%s", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Day00Part1(tt.in)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay00Part1(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(00))
	if err != nil {
		t.Fatal(err)
	}
	got := Day00Part1(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay00Part2(t *testing.T) {
}
