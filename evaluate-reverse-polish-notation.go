package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, t := range tokens {
		switch l := len(stack); t {
		case "+":
			// store the result of an operation in the second from last position and
			// slice only the last element from the stack
			stack[l-2] = stack[l-2] + stack[l-1]
			stack = stack[:l-1]
		case "-":
			stack[l-2] = stack[l-2] - stack[l-1]
			stack = stack[:l-1]
		case "*":
			stack[l-2] = stack[l-2] * stack[l-1]
			stack = stack[:l-1]
		case "/":
			stack[l-2] = stack[l-2] / stack[l-1]
			stack = stack[:l-1]
		default:
			i, _ := strconv.Atoi(t)
			stack = append(stack, i)
		}
	}
	return stack[0]
}

// Time: O(n)
// Space: O(n / 2 + 1) => O(n)
