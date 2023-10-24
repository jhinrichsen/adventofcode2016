package adventofcode2016

import (
	"fmt"
	"testing"
)

func TestDay20Part1Example(t *testing.T) {
	const (
		day   = 20
		part1 = true
		want  = 3
	)
	lines, err := linesFromFilename(exampleFilename(day))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day20(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// correct: f(5, 7) = 4
// correct: f(11, 14) = 8
func _TestDay20RangeMask(t *testing.T) {
	var want, lower, upper uint32
	want = 8
	lower = 11
	upper = 14
	got := rangeMask2(lower, upper)
	if want != got {
		t.Fatalf("want %0.8b but got %0.8b", want, got)
	}
}

func _TestDay20Lowest(t *testing.T) {
	var m15, m79, want uint32
	m15 |= 1
	m15 |= 2
	m15 |= 3
	m15 |= 4
	m15 |= 5
	want = 6
	m79 |= 7
	m79 |= 8
	m79 |= 9

	fmt.Printf("m15:   %08b\n", m15)
	fmt.Printf("m79:   %08b\n", m79)

	m1 := rangeMask3(1, 5)
	fmt.Printf("mask1: %08b\n", m1)
	if m15 != m1 {
		t.Fatalf("%d does not match %d", m15, m1)
	}
	m2 := rangeMask3(7, 9)
	fmt.Printf("mask2: %08b\n", m2)
	if m79 != m2 {
		t.Fatalf("%d does not match %d", m79, m2)
	}

	var got uint32 = m1 | m2
	fmt.Printf("want:  %08b\n", want)
	fmt.Printf("got:   %08b\n", got)
	if want != got {
		t.Fatalf("want %0.8b but got %0.8b", want, got)
	}
}

func _TestDay20RangeMask2(t *testing.T) {
	var want uint32
	// setup [0-4], [5-9]
	for i := want; i < 10; i++ {
		want |= i
	}
	fmt.Printf("want: %32b\n", want)
	got := rangeMask2(0, 9)
	fmt.Printf("got:  %32b\n", got)
	if want != got {
		t.Fatalf("want %b but got %b", want, got)
	}
}

func TestDay20Part1(t *testing.T) {
	const (
		day   = 20
		part1 = true
		want  = 23923783
	)
	lines, err := linesFromFilename(filename(day))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day20(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Part1(b *testing.B) {
	const (
		day   = 20
		part1 = true
	)
	lines, err := linesFromFilename(filename(day))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day20(lines, part1)
	}
}

func TestDay20Part2(t *testing.T) {
	const (
		day   = 20
		part1 = false
		want  = 125
	)
	lines, err := linesFromFilename(filename(day))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day20(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Part2(b *testing.B) {
	const (
		day   = 20
		part1 = false
	)
	lines, err := linesFromFilename(filename(day))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day20(lines, part1)
	}
}
