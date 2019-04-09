// Package leetcode :: leetcode.com
//		Ben Morrison
//		ben@gbmor.dev
package leetcode

// reverse() takes a signed integer as input, returning the same integer reversed.
// If the reversed integer will overflow a signed int32, returns 0.
//	On leetcode.com, ran in 4ms using 2.1MB memory, faster than 100%
// of all other submissions for this problem.
func reverse(x int) int {
	rev := 0
	if !math.Signbit(float64(x)) {
		rev = 0
		for x > 0 {
			remainder := x % 10
			rev *= 10
			rev += remainder
			x /= 10

		}

	} else {
		pos := int(math.Abs(float64(x)))
		ver := 0
		for pos > 0 {
			remainder := pos % 10
			ver *= 10
			ver += remainder
			pos /= 10

		}
		rev = -ver

	}
	if rev > int(math.Pow(2, 31)) {
		return 0

	} else if rev < -int(math.Pow(2, 31)) {
		return 0

	} else {
		return rev

	}

}
