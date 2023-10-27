package adventofcode2016

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func linesFromFilename(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func exampleFilename(day int) string {
	return fmt.Sprintf("testdata/day%d_example.txt", day)
}

func filename(day int) string {
	return fmt.Sprintf("testdata/day%d.txt", day)
}

// toint parses and returns a number as specifiec by strconv.Atoi or panic()s.
func toint(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
