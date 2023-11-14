package main

func coinChange(coins []int, amount int) int {
	// initialize an array where the index represents the amount up to target amount
	// and the value represents the minimum number of coins needed for each amount
	dp := make([]int, amount+1)

	for i := 1; i < amount+1; i++ {
		// initialize the current table value with amount + 1
		// this value will be updated with the minimum number of coins needed
		dp[i] = amount + 1

		for _, coin := range coins {
			if coin <= i {
				// for each coin that less or equal than current value
				// we calculate number of coins needed by
				// looking a previous minimum number of coins for
				// [current amount - current coin] + 1 (current coin)
				// and take minimum number of coins needed for that amount
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// if the initial value for the requested amount not updated
	// it indicates that there is no solution
	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

// Time: O(n)
// Space: O(n)
// Where n is the target amount + 1, and since it's maximum 12 coins, we consider it as a constant

// Helper function to find minimum between two integers
// Another option is to use math.Min()
// but it's more verbose and require conversion to and from float64
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
