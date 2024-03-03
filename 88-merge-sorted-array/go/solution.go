package solution

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
and two integers m and n, representing the number of elements in nums1 and nums2
respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function,
but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote
the elements that should be merged, and the last n elements are set to 0 and should
be ignored. nums2 has a length of n.
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	pointerNums1 := m - 1
	pointerNums2 := n - 1
	pointerMergedEnd := len(nums1) - 1

	for pointerNums2 >= 0 {
		if pointerNums1 >= 0 && (nums1[pointerNums1] >= nums2[pointerNums2]) {
			nums1[pointerMergedEnd] = nums1[pointerNums1]
			pointerNums1--
			pointerMergedEnd--
		} else {
			nums1[pointerMergedEnd] = nums2[pointerNums2]
			pointerNums2--
			pointerMergedEnd--
		}

	}
}
