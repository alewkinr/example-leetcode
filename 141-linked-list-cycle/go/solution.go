package solution

// ListNode — definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	TASK: Linked List Cycle
	Given head, the head of a linked list, determine if the linked list has a cycle in it.
	There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
	Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
	Note that pos is not passed as a parameter.

	Return true if there is a cycle in the linked list. Otherwise, return false.

	SOLUTION:
	There are several ways to solve the problem
	1 — using map, and caching visited nodes
	2 — transforming node structure, adding a visited flag
	3 — Floyd's Cycle Finding algo

	e.g. https://www.geeksforgeeks.org/detect-loop-in-a-linked-list/

	We use here the third method.
*/
// hasCycle — checks if Linked List has a cycle links
func hasCycle(head *ListNode) bool {
	/*
		There are several ways to solve the problem
		1 — using map, and caching visited nodes
		2 — transforming node structure, adding a visited flag
		3 — Floyd's Cycle Finding algo
		e.g. https://www.geeksforgeeks.org/detect-loop-in-a-linked-list/

	*/

	// PrintLinkedList(head)

	var result bool
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		// fast goes twice faster, so
		// slow will be in the middle of the list
		// when fast finishes
		fast = fast.Next.Next

		if fast == slow {
			result = true
			break
		}
	}

	return result
}
