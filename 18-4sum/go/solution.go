package solution

/*
	TASK: 4Sum

	Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

	0 <= a, b, c, d < n
	a, b, c, and d are distinct.
	nums[a] + nums[b] + nums[c] + nums[d] == target
	You may return the answer in any order.
*/

// fourSum — defines all the combination of numbers from nums which sum will be equal target
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	set := make(map[int]struct{}, 0)
	sumResult := make(map[int]map[int]int, 0)

	slow := 0
	fast := slow + 1

	for slow < len(nums)-1 {
		slowV := nums[slow]
		fastV := nums[fast]
		requiredSum := target - slowV - fastV

		if value, ok := sumResult[slowV+fastV]; ok {
			arr := make([]int, 0)
			for k, v := range value {
				if _, ok := set[v]; !ok {
					arr = append(arr, k)
					set[v] = struct{}{}
				}
			}
			arr = append(arr, slowV, fastV)
			result = append(result, arr)
		}

		sumResult[requiredSum] = make(map[int]int, 0)
		sumResult[requiredSum][slowV] = slow
		sumResult[requiredSum][fastV] = fast

		if fast < len(nums)-1 {
			fast++
		} else {
			slow++
			fast = slow + 1
		}

		// [0, 1,2,3,4,5,6,7], 10
		// [[0,2,3,5],...]
	}

	return result
}
