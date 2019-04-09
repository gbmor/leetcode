package leetcode

import (
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	var roman string
	var snum string

	if num < 10 {
		if num == 9 {
			return "IX"
		}
		for i := 0; i < num; i++ {
			roman += "I"

		}
		return roman

	}
	if (num / 1000) > 0 {
		for i := 0; i < (num / 1000); i++ {
			roman += "M"

		}

	}
	snum = strconv.Itoa(num / 100)

	if strings.HasSuffix(snum, "5") {
		roman += "D"

	}
	if strings.HasSuffix(snum, "4") {
		roman += "CD"

	}
	if strings.HasSuffix(snum, "9") {
		roman += "CM"

	}
	c := (num / 100 % 5)
	if c != 0 && c != 9 && c != 5 && c != 4 {
		for i := 0; i < (num / 100); i += 10 {
			roman += "C"

		}

	}
	snum = strconv.Itoa(num / 10)
	if strings.HasSuffix(snum, "5") {
		roman += "L"

	} else if strings.HasSuffix(snum, "4") {
		roman += "XL"

	}
	snum = strconv.Itoa(num)
	if strings.HasSuffix(snum, "15") {
		roman += "XV"

	} else if strings.HasSuffix(snum, "14") {
		roman += "XIV"

	} else if strings.HasSuffix(snum, "10") {
		roman += "X"

	}
	return roman

}
