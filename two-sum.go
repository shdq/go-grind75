package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}

	return nil
}

// Time: O(n)
// Space: O(n)
