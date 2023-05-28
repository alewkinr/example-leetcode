package main

import (
	"fmt"
	"time"
)

// ListNode — definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// PrintLinkedList — prints out a Linked List
func PrintLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%v\n", current)
		time.Sleep(time.Millisecond * 50)
		current = current.Next
	}
}

/***
	TASK: Delete Node in a Linked List

	There is a singly-linked list head and we want to delete a node node in it.

	You are given the node to be deleted node. You will not be given access to the first node of head.

	All the values of the linked list are unique, and it is guaranteed that the given node node is not the last node in the linked list.

	Delete the given node. Note that by deleting the node, we do not mean removing it from memory. We mean:

	The value of the given node should not exist in the linked list.
	The number of nodes in the linked list should decrease by one.
	All the values before node should be in the same order.
	All the values after node should be in the same order.
	Custom testing:

	For the input, you should provide the entire linked list head and the node to be given node. node should not be the last node of the list and should be an actual node in the list.
	We will build the linked list and pass the node to your function.
	The output will be the entire list after calling your function.
***/
// deleteNode — deletes node from linked list
func deleteNode(node *ListNode) {
	// node1 -> curNode -> node2
	// to be deleted, needs node1.Next change to node2

	// change value
	nodeRef := node
	nodeRef.Val = node.Next.Val
	nodeRef.Next = node.Next.Next
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

/*
	TASK:  Linked List Cycle

	Given head, the head of a linked list, determine if the linked list has a cycle in it.

	There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

	Return true if there is a cycle in the linked list. Otherwise, return false.
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

func main() {

	head := &ListNode{ // 0
		Val: 1,
		Next: &ListNode{ // 1
			Val: 2,
			Next: &ListNode{ // 2
				Val: 4,
				Next: &ListNode{ // 3
					Val: 4,
					Next: &ListNode{ // 4
						Val: 2,
						Next: &ListNode{ // 5
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}

	head2 := &ListNode{ // 0
		Val: 1,
		Next: &ListNode{ // 1
			Val: 3,
			Next: &ListNode{ // 2
				Val: 4,
				Next: &ListNode{ // 3
					Val: 25,
					Next: &ListNode{ // 4
						Val: 77,
						Next: &ListNode{ // 5
							Val:  93,
							Next: nil,
						},
					},
				},
			},
		},
	}

	head3 := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}

	node4 := &ListNode{Val: 4, Next: nil}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	node4.Next = node3

	head4 := &ListNode{Val: 123, Next: node1}

	// deleteNode(head)
	println(head, head2, head3)
	// removeNthFromEnd(head, 2)
	// removeNthFromEnd(head, 1)
	// removeNthFromEnd(&ListNode{Val: 1, Next: nil}, 1)
	// reverseList(head)
	// mergeTwoLists(head, head2)
	// println(isPalindrome(head3))
	println(hasCycle(head4))
}
