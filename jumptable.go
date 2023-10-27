package adventofcode2016

// jumpTable holds a scramble and a descramble jump table.
// descramble jump table may be incomplete for ambiguous command pipelines.
type jumpTable [2]map[int]int

// newJumpTable sets up a scramble/ descramble jump table for day 21.
// It is used for the 'rotate based on position' operation to allow an O(1) unscramble.
// For my example ("abcde"), no such jump table can be constructed, because f(b) = b' is not...
// not...
// not...
// well what is this math word if you can map b -> b' and b' <- b? in german it is 'eineindeutig'.
// unambigously transformable back and forth, 1:1.
// For the example "abcde", the jump table looks like
//
// Jump table for "abcde":
//
// b    delta       b'   b'%len
// 0        1        1        1
// 1        2        3        3
// 2        3        5        0    <-----
// 3        4        7        2
// 4        6       10        0    <----- ambiguous
//
// What this means is that the 'rotate based on position' cannot be inversed
// because b' == 0 can either result from b == 2 or b == 4.
//
// Jump table for "abcdefgh":
//
// b    delta       b'   b'%len
// 0        1        1        1
// 1        2        3        3
// 2        3        5        5
// 3        4        7        7
// 4        6       10        2
// 5        7       12        4
// 6        8       14        6
// 7        9       16        0
//
// If the jump table cannot be constructed for the given input, an empty map is returned.
// This can only happen for unscramble, read, scramble == false.
func newJumpTable(password string) jumpTable {
	f := func(x int) int {
		delta := 1
		delta += x
		if x >= 4 {
			delta++
		}
		return delta
	}
	// b -> delta
	bs := make(map[int]int)
	// b' -> delta
	bs_ := make(map[int]int)
	for b := 0; b < len(password); b++ {
		delta := f(b)
		bs[b] = delta
		b_ := (b + delta) % len(password)
		bs_[b_] = delta
	}
	return jumpTable{bs, bs_}
}
