package main

// ListNode is a node struct for a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// a dummy node to serve as the starting point of the merged list
	head := &ListNode{}
	curr := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	// if one of the lists is not fully processed, append the remaining nodes
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return head.Next
}
