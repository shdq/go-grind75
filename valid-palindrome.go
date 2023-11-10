package main

import "strings"

func isAlphaNum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if !isAlphaNum(str[l]) {
			l++
			continue
		}
		if !isAlphaNum(str[r]) {
			r--
			continue
		}
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// Time: O(n)
// Space: O(1)
