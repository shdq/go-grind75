package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := [26]int{}

	for i := range s {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

// Time: O(n) where n is the length of the string
// Space: O(1) because it's only 26 lowercase english letters
