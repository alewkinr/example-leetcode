package solution

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	r := 1
	prefixCount := len(strs)
	firstString := strs[0]

OUTER:
	for prefixCount == len(strs) {
		if r > len(firstString) {
			break
		}

		pr := firstString[:r]
		for _, s := range strs {
			if !strings.HasPrefix(s, pr) {
				prefixCount--
				break OUTER
			}
		}
		r++
	}
	return firstString[:r-1]
}
