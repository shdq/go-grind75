package main

// recursive solution with memoization
func climbStairsRecursive(n int) int {
	m := map[int]int{}

	var f func(int) int

	f = func(i int) int {
		if i <= 1 {
			return 1
		}
		if v, ok := m[i]; ok {
			return v
		}
		m[i] = f(i-1) + f(i-2)

		return m[i]
	}

	return f(n)
}

// Time: O(n)
// Space: O(n)
