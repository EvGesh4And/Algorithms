package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxSub := 0
	left := 0
	for right := range len(s) {
		m[s[right]]++
		for m[s[right]] > 1 {
			m[s[left]]--
			if m[s[left]] == 0 {
				delete(m, s[left])
			}
			left++
		}
		maxSub = max(maxSub, len(m))
	}
	return maxSub
}
