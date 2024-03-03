package solution

// ListNode — definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	TASK: Palindrome Linked List

	Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
*/
// isPalindrome — check if LinkedList is a palindome
func isPalindrome(head *ListNode) bool {
	// using two pointers method
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		// fast goes twice faster, so
		// slow will be in the middle of the list
		// when fast finishes
		fast = fast.Next.Next
	}

	// reversing potential palindrome list
	var prev *ListNode
	curr := slow
	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}

	// compare the first half and the reversed second half of the linked list
	p1 := head
	p2 := prev
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}
