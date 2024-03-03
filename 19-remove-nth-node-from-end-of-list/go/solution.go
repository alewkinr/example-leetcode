package solution

// ListNode — definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	TASK: Remove Nth Node From End of List

	Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/
// removeNthFromEnd — removes Nth node from linked list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := &ListNode{Next: head}
	fast := head
	result := slow
	for i := 0; fast != nil; i++ {
		fast = fast.Next
		// if iterator grater than n, we start to rearrange nodes, substituting slowPointerNode for the next one
		if i >= n {
			slow = slow.Next
		}
	}
	// covering case [1], 1
	slow.Next = slow.Next.Next

	return result.Next
}
