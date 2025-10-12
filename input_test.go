package adventofcode2016

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// Helper functions using testing.TB - only way to access files
func linesFromFilename(tb testing.TB, filename string) []string {
	tb.Helper()
	f, err := os.Open(filename)
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()

	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		tb.Fatal(err)
	}

	// Reset timer if this is a benchmark
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}

	return lines
}

func exampleFilename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", day)
}

func exampleNFilename(day uint8, n int) string {
	return fmt.Sprintf("testdata/day%02d_example%d.txt", day, n)
}

func example1Filename(day uint8) string {
	return exampleNFilename(day, 1)
}

func example2Filename(day uint8) string {
	return exampleNFilename(day, 2)
}

func example3Filename(day uint8) string {
	return exampleNFilename(day, 3)
}

func example4Filename(day uint8) string {
	return exampleNFilename(day, 4)
}

func filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
}

func file(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filename(day))
	if err != nil {
		tb.Fatal(err)
	}

	// Reset timer if this is a benchmark
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}

	return buf
}

func exampleFile(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(exampleFilename(day))
	if err != nil {
		tb.Fatal(err)
	}

	// Reset timer if this is a benchmark
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}

	return buf
}

