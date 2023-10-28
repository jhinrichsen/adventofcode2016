package adventofcode2016

import "testing"

func TestDay22(t *testing.T) {
	const (
		part1 = true
		// want  = 20636 // first try wrong, too high
		// i was using col 1 (size) and col 2 (used) instead of col 2 (used) and col 3 (avail)
		// does that count as 'off-by-one'? not sure...
		want = 937
	)
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day22(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay22(b *testing.B) {
	const (
		part1 = true
	)
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day22(lines, part1)
	}
}
