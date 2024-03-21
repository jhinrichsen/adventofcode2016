package adventofcode2016

import (
	"testing"
)

var day4Examples = []struct {
	in  string
	out bool
}{
	{"aaaaa-bbb-z-y-x-123[abxyz]", true},
	{"a-b-c-d-e-f-g-h-987[abcde]", true},
	{"not-a-real-room-404[oarel]", true},
	{"totally-real-room-200[decoy]", false},
}

func TestDay4Logic(t *testing.T) {
	for _, tt := range day4Examples {
		t.Run(tt.in, func(t *testing.T) {
			want := tt.out
			d, err := newDay4(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			got := d.real()
			if want != got {
				t.Fatalf("want %v but got %v", want, got)
			}
		})
	}
}

func TestDay4Example(t *testing.T) {
	const want = 1514
	var is []string
	for _, tt := range day4Examples {
		is = append(is, tt.in)
	}
	got, err := Day4Part1(is)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay4Part1(b *testing.B) {
	const want = 137896
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := Day4Part1(lines)
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}

func BenchmarkDay4Part2(b *testing.B) {
	const want = 501
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := Day4Part2(lines)
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}

func TestDay4Decrypt(t *testing.T) {
	const want = "very encrypted name"
	s := "qzmt-zixmtkozy-ivhz-343"
	got := decrypt(s)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}
