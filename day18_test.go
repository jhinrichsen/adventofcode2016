package adventofcode2016

import "testing"

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

func TestDay18LargeExample(t *testing.T) {
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
		got := next(lines[i-1])
		if want != got {
			t.Fatalf("want %q but got %q", want, got)
		}
	}

	got := Day18(lines[0], len(lines))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const (
		input = "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^"
		want  = 1974
	)
	got := Day18(input, 40)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	const (
		input = "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^"
		want  = 19991126
	)
	got := Day18(input, 400000)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay18Part2(b *testing.B) {
	const (
		input = "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^"
	)
	for i := 0; i < b.N; i++ {
		_ = Day18(input, 400000)
	}
}
