package adventofcode2016

import (
	"math"
	"math/bits"
)

// highestOneBitValue returns 16 for 19 (2^4).
func highestOneBitValue(n uint) uint {
	return 1 << (bits.UintSize - bits.LeadingZeros(n) - 1)
}

// Day19 calculates https://oeis.org/A032434.
// It had to be a sequence, no algorithm will run in a reasonable time
// for a number > 3_000_000.
// f(N) = 2L + 1 where N =2^M + L and 0 <= L < 2^M
// Look mom, O(1).
func Day19Part1(n uint) uint {
	return 2*(n-highestOneBitValue(n)) + 1
}

/*
def highest_power_of_3(n):
    option = 0
    while 3**option <= n:
        option +=1
    return 3 ** (option -1)

def answer(n):
    x = highest_power_of_3(n)
    if x == n:
        return x
    else:
        if n < 2 * x:
            return n%x
        else:
            return x + 2 * (n%x)
*/

func highestPowerOf3(n uint) uint {
	pow3 := func(n uint) uint {
		return uint(math.Pow(3, float64(n)))
	}
	var option uint
	for pow3(option) <= n {
		option++
	}
	return pow3(option - 1)
}

// https://oeis.org/A334473
func Day19Part2(n uint) uint {
	x := highestPowerOf3(n)
	if x == n {
		return x
	}
	if n < 2*x {
		return n % x
	}
	return x + 2*(n%x)
}
