// Package leetcode :: leetcode.com
// Ben Morrison
// ben@gbmor.dev
package leetcode

var a int
var b int
var c int

// fib() returns the sum of the fibonacci sequence
// up to a specified index.
// Leetcode.com - 0ms, 1.9MB
// Faster than 100%
func fib(N int) int {
	if N == 0 {
		return 0

	}
	if N == 1 {
		return 1

	}
	if N == 2 {
		return 1

	}
	a = 0
	b = 1
	c = 1
	for i := 2; i < N; i++ {
		a = b
		b = c
		c = a + b

	}
	return c

}
