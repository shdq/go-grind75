package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		// the idea is to do two steps for the fast pointer
		fast = fast.Next.Next
		// and one step for the slow
		slow = slow.Next
	}
	// the moment we reach the end of the list, slow pointer will be in the middle
	return slow
}
