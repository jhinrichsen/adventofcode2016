package adventofcode2016

const (
	safe = 46 // '.'
	trap = 94 // '^'
)

// derive uses ancestors to produce a new safe or trap.
// Its left and center tiles are traps, but its right tile is not.
// Its center and right tiles are traps, but its left tile is not.
// Only its left tile is a trap.
// Only its right tile is a trap.
func derive(left, right bool) bool {
	/* my working version
	b := (left && center && (!right)) ||
		(center && right && (!left)) ||
		(left && (!center) && (!right)) ||
		(right && (!left) && (!center))
	*/

	/* running the beast through a boolean algebra simplifier
	(a && b && (!c)) ||
	(b && c && (!a)) ||
	(a && (!b) && (!c)) ||
	(c && (!a) && (!b))

	=> (a && ~ c) || (~ a && c)
	*/

	// after a closer inspection, the expression
	// (left && (!right)) || ((!left) && right)
	// is a plain XOR

	// XOR the Go way:
	return left != right
}

func Day18(row string, count int) int {
	return Day18V3(row, count)
}
