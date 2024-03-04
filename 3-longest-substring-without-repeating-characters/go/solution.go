package solution

import "fmt"

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]struct{})
	maximum, l := 0, 0

	for r := 0; r < len(s); r++ {
		for isinset(set, s[r]) {
			del(&set, s[l])
			l++
		}

		set[s[r]] = struct{}{}
		if len(set) > maximum {
			maximum = len(set)
		}
	}

	for i, _ := range set {
		fmt.Println(string(i))
	}

	return maximum
}

func del(set *map[byte]struct{}, charCode byte) {
	delete(*set, charCode)
}

func isinset(set map[byte]struct{}, charCode byte) bool {
	_, ok := set[charCode]
	return ok
}
