package adventofcode2016

import (
	"strings"
	"testing"
)

func trapsAsString(bits []bool) string {
	var sb strings.Builder
	for i := 1; i < len(bits)-1; i++ {
		var c byte
		if bits[i] {
			c = trap
		} else {
			c = safe
		}
		sb.WriteByte(c)
	}
	return sb.String()
}

func TestTrapsAsString(t *testing.T) {
	want := "^^.^.^."
	got := trapsAsString([]bool{
		false, // left safe
		true, true, false, true, false, true, false,
		false, // left safe
	})
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay18ExampleV3(t *testing.T) {
	const want = 38
	got := Day18V3(day18Example[0], len(day18Example))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// TODO add Part1V3

func TestDay18Part2V3(t *testing.T) {
	got := Day18V3(input, repeat)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay18Part2V3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day18V3(input, repeat)
	}
}
