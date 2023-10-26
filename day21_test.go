// all string operations check for their corresponding NOP action,
// and take all possible shortcuts.

package adventofcode2016

import (
	"testing"
)

func TestDay21Example(t *testing.T) {
	lineResults := []string{
		"ebcda",
		"edcba",
		"abcde",
		"bcdea",
		"bdeac",
		"abdec",
		"ecabd",
		"decab",
	}
	lines, err := linesFromFilename(exampleFilename(21))
	if err != nil {
		t.Fatal(err)
	}

	got := "abcde"
	for i, line := range lines {
		lineno := i + 1
		f := parseDay21(line)

		want := lineResults[i]
		got = f(got)
		if want != got {
			t.Fatalf("line #%d: want %q but got %q", lineno, want, got)
		}
	}
}

func test(t *testing.T, name string, got, want string) {
	if want != got {
		t.Fatalf("error in %q: want %q but got %q", name, want, got)
	}
}

func TestSwapPosition(t *testing.T) {
	test(t, "swap position", swapPosition("abcde", 4, 0), "ebcda")
}

func TestSwapLetter(t *testing.T) {
	test(t, "swap letter", swapLetter("ebcda", "d", "b"), "edcba")
}

func TestReverse(t *testing.T) {
	test(t, "reverse", reverse("edcba", 0, 4), "abcde")
}

func TestRotateLeftN(t *testing.T) {
	s := "rotate left N"
	f := rotateLeftN
	test(t, s, f("abcde", 1), "bcdea")
	test(t, s, f("abcde", 4), "eabcd")
	test(t, s, f("abcde", 5), "abcde")
}

func TestRotateRightN(t *testing.T) {
	s := "rotate right N"
	f := rotateRightN
	test(t, s, f("abcde", 1), "eabcd")
	test(t, s, f("abcde", 4), "bcdea")
	test(t, s, f("abcde", 5), "abcde")
}

func TestRotateLeftPos(t *testing.T) {
	test(t, "rotate left pos", rotateLeftPos("abdec", "b"), "ecabd")
}

func TestRotateRightPos(t *testing.T) {
	test(t, "rotate right pos", rotateRightPos("ecabd", "d"), "decab")
}

func TestMove(t *testing.T) {
	test(t, "move", move("bcdea", 1, 4), "bdeac")
	test(t, "move", move("bdeac", 3, 0), "abdec")
}

func TestDay21(t *testing.T) {
	const want = "gfdhebac"
	lines, err := linesFromFilename(filename(21))
	if err != nil {
		t.Fatal(err)
	}
	got := Day21("abcdefgh", lines, true)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}
