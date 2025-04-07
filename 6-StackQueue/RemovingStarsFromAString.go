func removeStars(s string) string {
	stack := make([]rune, 0)

	for _, r := range s {
		if len(stack) > 0 && r == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}