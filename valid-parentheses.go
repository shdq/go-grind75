package main

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, r := range s {
		if v, ok := pairs[r]; ok {
			l := len(stack)
			if l == 0 {
				return false
			}
			pop := stack[l-1]
			stack = stack[:l-1]
			if pop != v {
				return false
			}
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}

// Time: O(n)
// Space: O(n)
