package adventofcode2016

import (
	"testing"
)

func TestDay18ExampleV2(t *testing.T) {
	const want = 38
	var lines = []string{
		".^^.^.^^^^",
		"^^^...^..^",
		"^.^^.^.^^.",
		"..^^...^^^",
		".^^^^.^^.^",
		"^^..^.^^..",
		"^^^^..^^^.",
		"^..^^^^.^^",
		".^^^..^.^^",
		"^^.^^^..^^",
	}

	for i := 1; i < len(lines); i++ {
		want := lines[i]
		from := newSafesAndTraps(lines[i-1])
		into := newSafesAndTraps(lines[i])
		step(from, into)
		got := into.String()
		if want != got {
			t.Fatalf("line %d:  want %q but got %q", i, want, got)
		}
	}

	got := Day18V2(lines[0], len(lines))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2V2(t *testing.T) {
	got := Day18V2(day18Input, day18RepeatPart2)
	if day18WantPart2 != got {
		t.Fatalf("want %d but got %d", day18WantPart2, got)
	}
}

func BenchmarkDay18Part2V2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day18V2(day18Input, day18RepeatPart2)
	}
}
