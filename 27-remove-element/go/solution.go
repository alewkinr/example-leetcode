package solution

func removeElement(nums []int, val int) int {
	var r, wr int

	for r < len(nums) {
		if nums[r] == val {
			r++
		} else {
			nums[wr] = nums[r]
			r++
			wr++
		}

	}

	return wr
}
