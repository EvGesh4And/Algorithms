package main

func generateParenthesis(n int) []string {
	var result []string

	var backtrack func(current string, open, close int)
	backtrack = func(current string, open, close int) {
		if len(current) == n*2 {
			result = append(result, current)
			return
		}

		if open < n {
			backtrack(current+"(", open+1, close)
		}
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}
