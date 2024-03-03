package solution

/*
	TASK: Single Number

	Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
	You must implement a solution with a linear runtime complexity and use only constant extra space.
*/
// singleNumber — find unique element in array with O(n) time and O(1) memory
func singleNumber(nums []int) int {
	unique := make(map[int]int, 0)
	var result int
	for i := 0; i < len(nums); i++ {
		unique[nums[i]] = unique[nums[i]] + 1

		if unique[nums[i]] == 1 {
			result += nums[i]
		} else {
			result -= nums[i]
		}
	}

	return result
}
