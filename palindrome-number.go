// Package leetcode :: leetcode.com
//		Ben Morrison
//		ben@gbmor.dev
package leetcode

// isPalindrome() takes an int as an argument,
// reverses the integer, and compares to the
// original. If the same, return true.
// Always returns false for negatives.
//	On leetcode.com, ran in 44ms with 5MB memory
// used. Faster than 100% of other submissions and
// less memory than 98.23% of others.
func isPalindrome(x int) bool {
	// just skip the crunch and return
	// false if it's negative
	if math.Signbit(float64(x)) {
		return false

	}
	// reverse the number
	rev := 0
	remainder := 0
	argz := x
	for argz > 0 {
		remainder = argz % 10
		rev *= 10
		rev += remainder
		argz /= 10

	}
	if rev == x {
		return true

	}
	return false

}
