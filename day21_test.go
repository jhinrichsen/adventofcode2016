// all string operations check for their corresponding NOP action,
// and take all possible shortcuts.

package adventofcode2016

import (
	"slices"
	"testing"
)

var day21ExampleResults = []string{
	"abcde", // start password
	"ebcda", // result after applying first command
	"edcba",
	"abcde",
	"bcdea",
	"bdeac",
	"abdec",
	"ecabd",
	"decab", // result
}

func diet(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func test(t *testing.T, name string, got, want string) {
	if want != got {
		t.Fatalf("error in %q: want %q but got %q", name, want, got)
	}
}

func testN(t *testing.T, cmds []string, results []string, scramble bool) {
	got := results[0]
	jt := newJumpTable(results[0])
	if !scramble && (len(jt[0]) != len(jt[1])) {
		t.Skip("ambiguous descramble")
	}

	for i, cmd := range cmds {
		f1, f2, err := compile(cmd, jt)
		if err != nil {
			t.Fatal(err)
		}
		var f stepfn
		if scramble {
			f = f1
		} else {
			f = f2
		}
		want := results[i+1]
		got = f(got)
		if want != got {
			idx := i + 1 // lines are 1-based
			t.Fatalf("line #%d: want %q but got %q", idx, want, got)
		}
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
	test(t, s, f("abcde", 0), "abcde")
	test(t, s, f("abcde", 1), "bcdea")
	test(t, s, f("abcde", 2), "cdeab")
	test(t, s, f("abcde", 3), "deabc")
	test(t, s, f("abcde", 4), "eabcd")
	test(t, s, f("abcde", 5), "abcde")
	test(t, s, f("abcde", 6), "bcdea")
}

func TestRotateRightN(t *testing.T) {
	s := "rotate right N"
	f := rotateRightN
	test(t, s, f("abcde", 0), "abcde")
	test(t, s, f("abcde", 1), "eabcd")
	test(t, s, f("abcde", 2), "deabc")
	test(t, s, f("abcde", 3), "cdeab")
	test(t, s, f("abcde", 4), "bcdea")
	test(t, s, f("abcde", 5), "abcde")
	test(t, s, f("abcde", 6), "eabcd")
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

func TestDay21Part1Example(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(21))
	testN(t, lines, day21ExampleResults, true)
}

func TestDay21Part2Example(t *testing.T) {
	cmds := linesFromFilename(t, exampleFilename(21))

	// reverse commands and expected results
	slices.Reverse(cmds)
	// create our own copy because tests may run in parallel
	results := make([]string, len(day21ExampleResults))
	copy(results, day21ExampleResults)
	slices.Reverse(results)

	testN(t, cmds, results, false)
}

func TestDay21Part1(t *testing.T) {
	const (
		input = "abcdefgh"
		part1 = true // part1 is synomym to 'scramble'
		want  = "gfdhebac"
	)
	lines := linesFromFilename(t, filename(21))
	got, err := Day21(lines, input, part1)
	diet(t, err)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay21Part1(b *testing.B) {
	const (
		input = "abcdefgh"
		part1 = true // part1 is synomym to 'scramble'
	)
	lines := linesFromFilename(b, filename(21))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day21(lines, input, part1)
	}
}

func TestDay21Part2(t *testing.T) {
	const (
		input = "fbgdceah"
		part1 = false // part1 is synomym to 'scramble'
		want  = "dhaegfbc"
	)
	cmds := linesFromFilename(t, filename(21))
	got, err := Day21(cmds, input, part1)
	diet(t, err)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay21Part2(b *testing.B) {
	const (
		input = "fbgdceah"
		part1 = false // part1 is synomym to 'scramble'
	)
	cmds := linesFromFilename(b, filename(21))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day21(cmds, input, part1)
	}
}
