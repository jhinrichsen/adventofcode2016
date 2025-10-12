package adventofcode2016

import (
	"testing"
)

func TestDay01Examples(t *testing.T) {
	tests := []struct {
		in    string
		want  uint
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

	for _, tt := range tests {
		name := tt.in
		if tt.part2 {
			name += " (part 2)"
		}
		t.Run(name, func(t *testing.T) {
			got := Day01(tt.in, !tt.part2)
			if got != tt.want {
				t.Errorf("Day01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay01Part1(t *testing.T) {
	const want = 299
	lines := linesFromFilename(t, filename(1))
	got := Day01(lines[0], true)
	if got != want {
		t.Errorf("Day01() = %v, want %v", got, want)
	}
}

func TestDay01Part2(t *testing.T) {
	const want = 181
	lines := linesFromFilename(t, filename(1))
	got := Day01(lines[0], false)
	if got != want {
		t.Errorf("Day01() = %v, want %v", got, want)
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(1))
	for b.Loop() {
		Day01(lines[0], true)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(1))
	for b.Loop() {
		Day01(lines[0], false)
	}
}
