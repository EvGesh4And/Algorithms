package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	left := 0
	res := 0
	for i := range len(s) {
		m[s[i]]++
		if m[s[i]] > 1 {
			for s[left] != s[i] {
				delete(m, s[left])
				left++
			}
			m[s[left]]--
			left++
		}
		res = max(res, len(m))
	}
	return res
}
