// Package leetcode :: leetcode.com
//		Ben Morrison
//		ben@gbmor.dev
package leetcode

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// romanToInt() takes a string Roman Numeral as input
// and converts is to a decimal number.
//		On leetcode.com, this ran in 16ms using 3MB memory,
// which was reported as faster and less memory use than
// 100% of submissions for this exercise.
func romanToInt(s string) int {
	var num int
	for i, e := range []byte(s) {
		num += roman[string(e)]
		if i > 0 {
			if (e == 'X' || e == 'V') && s[i-1] == 'I' {
				num -= 2

			}
			if (e == 'L' || e == 'C') && s[i-1] == 'X' {
				num -= 20

			}
			if (e == 'D' || e == 'M') && s[i-1] == 'C' {
				num -= 200

			}

		}

	}
	return num

}
