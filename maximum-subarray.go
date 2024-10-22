package main

// min and max functions were added in go 1.21

func maxSubArray(nums []int) int {
	// maximum sum and current sum
	m, s := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		// we choose which is more, current sum up until now or current element
		s = max(s+nums[i], nums[i])
		// update maximum
		m = max(m, s)
	}

	return m
}

// Time: O(n)
// Space: O(1)
