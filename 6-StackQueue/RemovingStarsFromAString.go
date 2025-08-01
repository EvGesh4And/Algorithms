// Будто это правильнее, так как в примере ниже если len(stack) == 0, то добавляем *
func removeStars(s string) string {
	var stack []rune
	for _, r := range s {
		if r != '*' {
			stack = append(stack, r)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

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