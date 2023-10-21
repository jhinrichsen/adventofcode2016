package adventofcode2016

import (
	"testing"
)

func TestDay1Examples(t *testing.T) {
	var tableTests = []struct {
		in    string
		out   uint
		part2 bool
	}{
		{"R2, L3", 5, false},
		{"R2, R2, R2", 2, false},
		{"R5, L5, R5, R3", 12, false},

		//     ^
		//     o
		//     |
		//     |
		// ^   |
		// o---X---o>
		//     |   |
		//     |   |
		//     |   |
		//    <o---o
		//         v
		{"R8, R4, R4, R8", 4, true},
	}

	for _, tt := range tableTests {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := day1(tt.in, tt.part2)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay1Part1(b *testing.B) {
	const want = 299
	lines, err := linesFromFilename(filename(1))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := Day1Part1(lines[0])
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}

func BenchmarkDay1Part2(b *testing.B) {
	const want = 181
	lines, err := linesFromFilename(filename(1))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := Day1Part2(lines[0])
		if err != nil {
			b.Fatal(err)
		}
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
