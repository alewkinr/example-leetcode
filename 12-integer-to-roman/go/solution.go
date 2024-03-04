package solution

func intToRoman(num int) string {
	bases := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var romanSet = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	var res string
	for _, base := range bases {
		baseCount := div(num, base)

		// num has current base
		if baseCount > 0 {
			res += printRoman(romanSet[base], baseCount)
			num -= base * baseCount
		}
		// exit if num is low
		if num <= 0 {
			break
		}
	}
	return res

}

func div(num int, base int) int {
	return num / base
}

func printRoman(symb string, times int) string {
	var res string
	for i := 1; i <= times; i++ {
		res += symb
	}
	return res
}
