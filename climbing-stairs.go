package main

// basically this is a fibonacci problem
func climbStairs(n int) int {
	n1, n2 := 0, 1

	for i := 0; i < n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}

// Time: O(n)
// Space: O(1)
