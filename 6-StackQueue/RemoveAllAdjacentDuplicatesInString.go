func removeDublicates(s string) string {
	var stack []rune

	for _, r := range s {
		if len(stack) == 0 || stack[len(stack)-1] != r {
			stack = append(stack, r)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

func removeDuplicates(s string) string {
	stack := make([]rune, 0)
	for _, v := range s {
		if len(stack) > 0 && stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}