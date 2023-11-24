package main

func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}

	// frequency count
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}

	var ans int
	for _, v := range m {
		// smarter way to get even amout of chars without using modulo % operator
		// for example, for the string 'aaaaa', where the frequency is 5, (5 / 2) * 2 = 4.
		ans += v / 2 * 2
	}

	// if the calculated length is not equal to the length of the original string,
	// it means there was an odd frequency of letters, so increment the answer by 1,
	// as it allows placing one character in the middle of the palindrome
	if ans < len(s) {
		ans += 1
	}

	return ans
}

// Time: O(n)
// Space: O(1), the maximum map size is 52 which is constant
