// Package leetcode :: leetcode.com
// Ben Morrison
// ben@gbmor.dev
package leetcode

// sortArray takes an array of ints, sorts it with even
// numbers at the beginning, then returns it.
// On leetcode.com: 52ms run time, 8.2MB memory use
// Faster than 100%, but memory use less than only 44.44%,
// of submissions
func sortArrayByParity(A []int) []int {
	var one []int
	var two []int
	for _, e := range A {
		if e%2 == 0 {
			one = append(one, e)

		} else {
			two = append(two, e)

		}

	}
	return append(one, two...)

}
