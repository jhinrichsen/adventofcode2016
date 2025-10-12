package adventofcode2016

import (
	"testing"
)

var day2Examples = []struct {
	id    string
	part1 bool
	out   string
}{
	{"part 1", true, "1985"},
	{"part 2", false, "5DB3"},
}

func TestDay2Part1Example(t *testing.T) {
	for _, tt := range day2Examples {
		lines := linesFromFilename(t, exampleFilename(2))
		want := tt.out
		got, err := Day2(lines, tt.part1)
		if err != nil {
			t.Fatal(err)
		}
		if want != got {
			t.Fatalf("want %q but got %q", want, got)
		}
	}
}

func TestDay02Part1(t *testing.T) {
	const want = "65556"
	lines := linesFromFilename(t, filename(2))
	got, err := Day2(lines, true)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Day2() = %v, want %v", got, want)
	}
}

func TestDay02Part2(t *testing.T) {
	const want = "CB779"
	lines := linesFromFilename(t, filename(2))
	got, err := Day2(lines, false)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Day2() = %v, want %v", got, want)
	}
}

func BenchmarkDay2Part1(b *testing.B) {
	benchDay2(b, "65556", true)
}

func BenchmarkDay2Part2(b *testing.B) {
	benchDay2(b, "CB779", false)
}

func benchDay2(b *testing.B, want string, part1 bool) {
	lines := linesFromFilename(b, filename(2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := Day2(lines, part1)
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %q but got %q", want, got)
		}
	}
}
