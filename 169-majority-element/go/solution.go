package solution

// majorityElement
func majorityElement(nums []int) int {
	var counter, res int

	for _, v := range nums {
		if counter == 0 {
			res = v
		}

		if v == res {
			counter++
		} else {
			counter--
		}

	}

	return res
}
