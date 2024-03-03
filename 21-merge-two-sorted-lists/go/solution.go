package solution

// ListNode — definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	TASK: Merge Two Sorted Lists

	You are given the heads of two sorted linked lists list1 and list2.

	Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

	Return the head of the merged linked list.

*/
// mergeTwoLists — merge to linked lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current1, current2 := list1, list2
	resultHead := new(ListNode)
	resultCurrent := resultHead

	for current1 != nil && current2 != nil {

		// compare 1
		if current1.Val <= current2.Val {
			resultCurrent.Next = current1
			current1 = current1.Next
		} else {
			// else 'coz each loop step we need only one comparance
			resultCurrent.Next = current2
			current2 = current2.Next
		}
		// shift resulting list insert pos
		resultCurrent = resultCurrent.Next

	}
	// merge the rest of list1/list2 as a tail of result lis
	if current1 != nil {
		resultCurrent.Next = current1
	} else {
		resultCurrent.Next = current2
	}

	// PrintLinkedList(resultHead.Next)
	return resultHead.Next
}
