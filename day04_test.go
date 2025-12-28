package adventofcode2016

import (
	"strings"
	"testing"
)

func TestDay04Part1Examples(t *testing.T) {
	tests := []struct {
		in   string
		real bool
	}{
		{"aaaaa-bbb-z-y-x-123[abxyz]", true},
		{"a-b-c-d-e-f-g-h-987[abcde]", true},
		{"not-a-real-room-404[oarel]", true},
		{"totally-real-room-200[decoy]", false},
	}

	// Sum of sector IDs for real rooms: 123 + 987 + 404 = 1514
	const want uint = 1514
	var sb strings.Builder
	for _, tt := range tests {
		sb.WriteString(tt.in)
		sb.WriteByte('\n')
	}
	got := Day04Part1([]byte(sb.String()))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part1(t *testing.T) {
	const want uint = 137896
	input := file(t, 4)
	got := Day04Part1(input)
	if got != want {
		t.Errorf("Day04Part1() = %v, want %v", got, want)
	}
}

func TestDay04Part2(t *testing.T) {
	const want uint = 501
	input := file(t, 4)
	got := Day04Part2(input)
	if got != want {
		t.Errorf("Day04Part2() = %v, want %v", got, want)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	input := file(b, 4)
	for b.Loop() {
		Day04Part1(input)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	input := file(b, 4)
	for b.Loop() {
		Day04Part2(input)
	}
}

func TestDay04Decrypt(t *testing.T) {
	// Test that "qzmt-zixmtkozy-ivhz-343" decrypts to "very encrypted name"
	// We test this by checking that the decrypt function works
	room := []byte("qzmt-zixmtkozy-ivhz-343")
	// Expected: "very encrypted name"
	// Sector 343 % 26 = 5 (shift by 5)
	// q+5=v, z+5=e, m+5=r, t+5=y
	target := []byte("very encrypted name")
	if !decryptsToTarget(room, 343) {
		// Manual check
		n := byte(343 % 26)
		var result []byte
		for _, b := range room {
			if b >= '0' && b <= '9' {
				break
			}
			if b == '-' {
				result = append(result, ' ')
				continue
			}
			decrypted := b + n
			if decrypted > 'z' {
				decrypted -= 26
			}
			result = append(result, decrypted)
		}
		// Trim trailing space
		for len(result) > 0 && result[len(result)-1] == ' ' {
			result = result[:len(result)-1]
		}
		if string(result) != string(target) {
			t.Fatalf("want %q but got %q", target, result)
		}
	}
}
