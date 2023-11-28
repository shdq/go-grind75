package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		// move slow pointer with one step and fast pointer with two steps
		slow = slow.Next
		fast = fast.Next.Next
		// if pointers point to the same node, there is a cycle
		if slow == fast {
			return true
		}
	}

	return false
}

// Time: O(n)
// Space: O(1) – Floyd's Tortoise and Hare Algorithm allow us to avoid using extra memory
