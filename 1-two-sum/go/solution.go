package solution

/*
	TASK: Two Sum

	Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.
	You can return the answer in any order.
*/
// twoSum — returns two numbers which sum will be equal to target
func twoSum(nums []int, target int) []int {
	sumResult := make(map[int]int, 0)
	var result [2]int

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		requiredNumber := target - cur

		// check if we already have pair for current number
		if value, ok := sumResult[cur]; ok {
			// if we do, return current and previous numbers as result
			result[0] = i
			result[1] = value
			break
		}
		// if we don't, fill required pair (number, which we need to fit the target sum for current number)
		// and current value into map
		sumResult[requiredNumber] = i
	}
	return result[:]
}
