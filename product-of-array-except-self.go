package main

func productExceptSelf(nums []int) []int {
	var prefix, postfix []int

	prefix = append(prefix, nums[0])
	for i := 1; i < len(nums); i++ {
		prefix = append(prefix, nums[i]*prefix[i-1])
	}

	postfix = append(postfix, nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		postfix = append(postfix, nums[i]*postfix[len(postfix)-1])
	}

	reverse(postfix)

	ans := []int{}

	for i := range nums {
		if i == 0 {
			ans = append(ans, postfix[i+1])
			continue
		}
		if i == len(nums)-1 {
			ans = append(ans, prefix[i-1])
			continue
		}
		ans = append(ans, prefix[i-1]*postfix[i+1])
	}

	return ans
}

// in place slice reverse function
func reverse(s []int) {
	l := len(s) - 1
	for i := 0; i <= l/2; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}
}

// Time: O(n)
// Space: O(n)
