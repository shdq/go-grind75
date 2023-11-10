func canConstruct(ransomNote string, magazine string) bool {
	m := [26]int{}

	for _, r := range magazine {
		m[r-'a']++
	}

	for _, r := range ransomNote {
		if m[r-'a'] == 0 {
			return false
		}
		m[r-'a']--
	}

	return true
}

// Time: O(n+m)
// Space: O(1)
