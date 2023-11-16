package main

func search(nums []int, target int) int {
	var lo, hi int = 0, len(nums) - 1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if nums[lo] == target {
		return lo
	}
	return -1
}

// Time: O(log n)
// Space: O(1)
