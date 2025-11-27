package adventofcode2016

import (
	"strings"
	"testing"
)

func TestDay13Example(t *testing.T) {
	const want = 11
	got, err := Day13Part1(10, 1+1i, 7+4i)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13ExampleString(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(13))
	want := strings.Join(lines, "\n")

	d := newDay13(10)
	got := d.String()
	if want != got {
		t.Fatalf("want:\n%s\nbut got\n%s\n", want, got)
	}
}

func TestDay13Part1(t *testing.T) {
	const (
		input = 1362
		want  = 82
	)
	got, err := Day13Part1(input, 1+1i, 31+39i)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay13Part1(b *testing.B) {
	for b.Loop() {
		_, _ = Day13Part1(1362, 1+1i, 31+39i)
	}
}

func TestDay13Part2(t *testing.T) {
	const (
		input = 1362
		want  = 138
	)
	got := Day13Part2(input, 1+1i, 50)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay13Part2(b *testing.B) {
	for b.Loop() {
		Day13Part2(1362, 1+1i, 50)
	}
}
