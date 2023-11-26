package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// initialize a pointer as nil
	var prev *ListNode
	curr := head

	for curr != nil {
		// save the next node of the current node
		next := curr.Next
		// reverse the direction of the current node to the previous node
		curr.Next = prev
		// make current node previous and move to the next node
		prev = curr
		curr = next
	}
	// the prev pointer now points to the head of the reversed list
	return prev
}
