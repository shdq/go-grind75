package main

func productExceptSelf(nums []int) []int {
	ans := []int{}

	prefix := 1
	for i := 0; i < len(nums)-1; i++ {
		if i == 0 {
			ans = append(ans, prefix)
		}
		prefix *= nums[i]
		ans = append(ans, prefix)
	}

	prefix = 1
	for i := len(nums) - 1; i > 0; i-- {
		prefix *= nums[i]
		ans[i-1] *= prefix
	}

	return ans
}

// Time: O(n)
// Space: O(1) output array isn't considered as extra memory in this problem
