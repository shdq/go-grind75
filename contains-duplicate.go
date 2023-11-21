package main

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}

	for _, v := range nums {
		if ok := m[v]; ok {
			return true
		}
		m[v] = true
	}

	return false
}

// Time: O(n)
// Space: O(n)

/*
  If we cannot afford extra space,
  an alternative solution involves O(n log n) time complexity
  and O(1) space complexity.
  This can be achieved by sorting the nums array and checking adjacent elements.
*/
