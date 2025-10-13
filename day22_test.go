package adventofcode2016

import "testing"

func TestDay22Part1(t *testing.T) {
	const (
		part1 = true
		// want  = 20636 // first try wrong, too high
		// i was using col 1 (size) and col 2 (used) instead of col 2 (used) and col 3 (avail)
		// does that count as 'off-by-one'? not sure...
		want = 937
	)
	lines := linesFromFilename(t, filename(22))
	got, err := Day22(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part2(t *testing.T) {
	const (
		part1 = false
		// want  = 181 // first try wrong, too low
		// want = 187 // fixing dim+1, too low
		want = 188 // counting moves instead of images
	)
	lines := linesFromFilename(t, filename(22))
	got, err := Day22(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay22Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(22))
	for b.Loop() {
		Day22(lines, true)
	}
}

func BenchmarkDay22Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(22))
	for b.Loop() {
		Day22(lines, false)
	}
}
