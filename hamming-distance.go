// Package leetcode :: leetcode.com
// Ben Morrison
// ben@gbmor.dev
package leetcode

import (
	"math/bits"
)

// hammingDistance() returns the 'hamming distance'
// of two integers.
// Leetcode.com stats: 0ms, 2MB memory.
// Ahead of 100% of submissions.
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x) ^ uint(y))

}
