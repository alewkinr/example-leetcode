package solution

import (
	"fmt"
	"strconv"
)

/*
	Task: Reverse Integer

	Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
	Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/
// reverse — reverse Integer
func reverse(x int) int {
	var isNegative bool
	if x < 0 {
		isNegative, x = true, x*-1
	}

	s := []byte(fmt.Sprintf("%v", x))

	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]

		left++
		right--
	}

	res, err := strconv.ParseInt(string(s), 10, 32)
	if err != nil {
		return 0
	}
	if isNegative {
		return int(res) * -1
	}
	return int(res)
}
