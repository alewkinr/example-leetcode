package solution

/*

	TASK: Reverse String

	Write a function that reverses a string. The input string is given as an array of characters s.
	You must do this by modifying the input array in-place with O(1) extra memory.
*/
// reverseString — reverses a string
func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]

		left++
		right--
	}
}
