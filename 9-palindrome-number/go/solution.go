package solution

import "strconv"

func isPalindrome(x int) bool {
	orig := strconv.Itoa(x)
	reversed := reverse(x)

	if orig == reversed {
		return true
	}
	return false

}

func reverse(x int) string {
	var res, sign string

	if x < 0 {
		x *= -1
		sign = "-"
	}

	if x == 0 {
		return "0"
	}

	for x != 0 {
		num := x % 10
		x /= 10
		res = res + strconv.Itoa(num)
	}

	return res + sign
}
