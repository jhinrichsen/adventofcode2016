package adventofcode2016

import (
	"fmt"
	"testing"
)

const (
	samplePasscode = "hijkl"
	sampleHashcode = "ced9"
)

func TestMD5(t *testing.T) {
	want := sampleHashcode
	got := md5lead([]byte(samplePasscode))
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestOpenDoor(t *testing.T) {
	open := openDoors([]byte(sampleHashcode))

	// all doors are open except for right (9)
	if len(open) != 3 {
		t.Fatalf("want %d open doors but got %d", 3, len(open))
	}
}

func TestWall(t *testing.T) {
	var tests = []struct {
		in  Position
		out bool // wall?
	}{
		{-1 + 0i, true},
		{0 - 1i, true},
		{0 + 4i, true},
		{4 + 0i, true},

		{0 + 0i, false}, // bottom left
		{3 + 3i, false}, // top right

		{StartPosition, false}, // top left
		{EndPosition, false},   // bottom right
	}
	for _, tt := range tests {
		id := fmt.Sprintf("%v", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := wall(tt.in)
			if want != got {
				t.Fatalf("%v: want %v but got %v", tt.in, want, got)
			}
		})
	}
}
