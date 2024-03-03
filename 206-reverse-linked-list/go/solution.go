package solution

// ListNode — definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	TASK: Reverse Linked List

	Given the head of a singly linked list, reverse the list, and return the reversed list.
*/
// reverseList — reverses a linked list
func reverseList(head *ListNode) *ListNode {
	current := head
	previous := (&ListNode{}).Next
	for current != nil {
		next := current.Next

		current.Next = previous
		previous = current
		current = next
	}

	head = previous
	// PrintLinkedList(head)
	return head
}
