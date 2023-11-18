package main

import (
	"math"
	"strings"
)

func addBinary(a string, b string) string {
	var ans []byte

	// get length difference
	diff := int(math.Abs(float64(len(a) - len(b))))

	// pad start smaller string with zeroes
	if len(b) > len(a) {
		a = strings.Repeat("0", diff) + a
	} else {
		b = strings.Repeat("0", diff) + b
	}

	carry := 0
	// going backwards for binary addition
	for i := len(a) - 1; i >= 0; i-- {
		// +'0' and -'0' here for proper conversion because of ASCII representation
		// calc the sum of the corresponding bits in a, b, carry
		sum := int(a[i]-'0') + int(b[i]-'0') + carry
		// append the least significant bit of the sum to the result
		ans = append(ans, byte(sum%2+'0'))
		// determine carry for the next higher bit position in binary addition
		carry = sum / 2
	}

	if carry == 1 {
		ans = append(ans, byte(carry+'0'))
	}

	// reverse string
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return string(ans)
}

// Time: O(n)
// Space: O(n) because of creating ans variable

/*
	Decision regarding whether ans should be a string or a byte slice:
	append to a byte slice + reverse in place is more efficient
	than concatenating strings using the + operator,
	because strings are immutable and concatenating creates a new string
*/
