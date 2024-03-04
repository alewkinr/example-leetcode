package solution

import "fmt"

// ListNode — Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := make([]int, 0)

	cur1, cur2 := l1, l2
	for i := 0; i <= 100; i++ {
		res := value(cur1) + value(cur2)
		tmp = append(tmp, res)

		cur1, cur2 = next(cur1), next(cur2)
		if cur1 == nil && cur2 == nil {
			break
		}
	}

	return buildNodeList(tmp)
}

func buildNodeList(s []int) *ListNode {
	var startNode *ListNode

	fmt.Println(s)

	nextNode := &ListNode{}
	curNode := &ListNode{}
	for i, v := range s {
		curNode = fill(curNode, v, nextNode, i, s)

		if i == 0 {
			startNode = curNode
		}

		curNode, nextNode = nextNode, &ListNode{}
		if i == len(s)-1 {
			break
		}
	}

	return startNode
}

func fill(
	node *ListNode,
	value int,
	next *ListNode,
	listIdx int,
	list []int,
) *ListNode {

	value = node.Val + value

	if value >= 10 {
		value %= 10
		next.Val += 1
	}

	if listIdx == len(list)-1 && next.Val == 0 {
		node.Val, node.Next = value, nil
		return node
	}

	if listIdx == len(list)-1 && next.Val > 0 {
		node.Val, node.Next = value, next
		return node
	}

	node.Val, node.Next = value, next
	return node
}

func next(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}

func value(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}
