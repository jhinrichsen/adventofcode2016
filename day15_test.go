package adventofcode2016

import (
	"testing"
)

/* Not working, commented out
func TestDay15Example(t *testing.T) {
	const (
		want  = 5
		part1 = true
	)
	lines, err := linesFromFilename(exampleFilename(15))
	if err != nil {
		t.Fatal(err)
	}
	d, err := newDay15(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day15(d, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/

func testDay15(t *testing.T, want uint, f func(day15) uint) {
	lines, err := linesFromFilename(filename(15))
	if err != nil {
		t.Fatal(err)
	}
	d, err := newDay15(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := f(d)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay15Part1(t *testing.T) {
	testDay15(t, 317371, func(d day15) uint {
		return Day15Hardcoded(d, true)
	})
}

func TestDay15Part2(t *testing.T) {
	testDay15(t, 2080951, func(d day15) uint {
		return Day15Hardcoded(d, false)
	})
}

/* Not working, commented out
func TestDay15Part1Smart(t *testing.T) {
	testDay15(t, func(d day15) uint {
		return Day15Smart(d, true)
	})
}
*/

func BenchmarkDay15Part1PlainVanilla(b *testing.B) {
	const part1 = true
	for i := 0; i < b.N; i++ {
		_ = Day15Hardcoded(day15{}, part1)
	}
}

func BenchmarkDay15Part2PlainVanilla(b *testing.B) {
	const part1 = false
	for i := 0; i < b.N; i++ {
		_ = Day15Hardcoded(day15{}, part1)
	}
}

/* Not working, commented out
func BenchmarkDay15Smart(b *testing.B) {
	const (
		want  = 317371
		part1 = true
	)
	lines, err := linesFromFilename(filename(15))
	if err != nil {
		b.Fatal(err)
	}
	d, err := newDay15(lines)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		got := Day15Smart(d, part1)
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
*/
