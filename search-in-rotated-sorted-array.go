package main

func searchInRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// if left side isn't rotated (sorted in ascending order)
		if nums[mid] >= nums[left] {
			// if the target is in the range of the left side
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				// overwise, take the right side
				left = mid + 1
			}

		} else { // the right side isn't rotated
			// if the target is in the range of the right side
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				// overwise, take the left side
				right = mid - 1
			}
		}
	}

	return -1
}

// Time: O(log n)
// Space: O(1)
