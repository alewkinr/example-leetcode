package solution

var romanSet = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var sum int
	for i := 0; i < len(s); i++ {
		r := string(s[i])

		if isException(r) {
			if i != len(s)-1 {
				v, ok := processSubstractions(r, string(s[i+1]))
				sum += v
				if ok {
					i++
				}
				continue
			}
		}
		sum += romanSet[r]
	}

	return sum
}

func isException(char string) bool {
	v, _ := romanSet[char]
	if (v == 1 || v%10 == 0) && v <= 100 {
		return true
	}
	return false
}

func processSubstractions(cur, next string) (int, bool) {
	cv, _ := romanSet[cur]
	nv, _ := romanSet[next]

	if (nv/5 == 1 ||
		nv/5 == 2 ||
		nv/5 == 10 ||
		nv/5 == 20 ||
		nv/5 == 100 ||
		nv/5 == 200) && nv > cv {
		return nv - cv, true
	}
	return cv, false

}
