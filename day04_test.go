package adventofcode2016

import (
	"testing"
)

func TestNewDay04Examples(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"aaaaa-bbb-z-y-x-123[abxyz]", true},
		{"a-b-c-d-e-f-g-h-987[abcde]", true},
		{"not-a-real-room-404[oarel]", true},
		{"totally-real-room-200[decoy]", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			want := tt.out
			d, err := newDay04(tt.in)
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

func TestDay04Part1Examples(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"aaaaa-bbb-z-y-x-123[abxyz]", true},
		{"a-b-c-d-e-f-g-h-987[abcde]", true},
		{"not-a-real-room-404[oarel]", true},
		{"totally-real-room-200[decoy]", false},
	}

	const want = 1514
	var is []string
	for _, tt := range tests {
		is = append(is, tt.in)
	}
	got, err := Day04Part1(is)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part1(t *testing.T) {
	const want = 137896
	lines := linesFromFilename(t, filename(4))
	got, err := Day04Part1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Day04Part1() = %v, want %v", got, want)
	}
}

func TestDay04Part2(t *testing.T) {
	const want = 501
	lines := linesFromFilename(t, filename(4))
	got, err := Day04Part2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Day04Part2() = %v, want %v", got, want)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(4))
	for b.Loop() {
		_, _ = Day04Part1(lines)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(4))
	for b.Loop() {
		_, _ = Day04Part2(lines)
	}
}

func TestDay04Decrypt(t *testing.T) {
	const want = "very encrypted name"
	s := "qzmt-zixmtkozy-ivhz-343"
	got := decrypt(s)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}
