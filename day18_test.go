package adventofcode2016

import (
	"reflect"
	"testing"
)

const (
	input  = "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^"
	repeat = 400000
	want   = 19991126
)

var day18Example = []string{
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

func TestDay18Example(t *testing.T) {
	const (
		in   = "..^^."
		want = ".^^^^"
	)
	got := next(in)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay18NewSafesAndTraps(t *testing.T) {
	s := "^.^.^.^^^^"
	want := safesAndTraps{
		false, // padding
		true, false, true, false, true, false, true, true, true, true,
		false} // padding

	got := newSafesAndTraps(s)
	if len(want) != len(got) {
		t.Fatalf("want len %d but got %d", len(want), len(got))
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", []bool(want), []bool(got))
	}

}

func TestDay18String(t *testing.T) {
	bits := safesAndTraps{
		false, // padding
		false, true, true, false, true, false, true, true, true, true,
		false} // padding
	want := ".^^.^.^^^^"
	got := bits.String()
	if len(want) != len(got) {
		t.Fatalf("want len %d but got %d", len(want), len(got))
	}
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}

}

// TestDay18Conversion tests one string -> bits -> string transformation to
// make sure the safe padding is correct.
func TestDay18Conversion(t *testing.T) {
	wants := []string{
		".^^.^.^^^^",
		input,
	}
	for _, want := range wants {
		got := newSafesAndTraps(want).String()
		if len(want) != len(got) {
			t.Fatalf("want len %d but got %d", len(want), len(got))
		}
		if want != got {
			t.Fatalf("want %q but got %q", want, got)
		}
	}
}

func TestDay18ExampleV1(t *testing.T) {
	const want = 38
	for i := 1; i < len(day18Example); i++ {
		want := day18Example[i]
		got := next(day18Example[i-1])
		if want != got {
			t.Fatalf("want %q but got %q", want, got)
		}
	}

	got := Day18V1(day18Example[0], len(day18Example))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const want = 1974
	got := Day18(input, 40)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	TestDay18Part2V1(t)
}

func TestDay18Part2V1(t *testing.T) {
	got := Day18(input, repeat)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay18Part2V1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day18V1(input, repeat)
	}
}

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
	got := Day18V2(input, repeat)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay18Part2V2(b *testing.B) {
	const (
		input = "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^"
	)
	for i := 0; i < b.N; i++ {
		_ = Day18V2(input, repeat)
	}
}
