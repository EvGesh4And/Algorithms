package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/dd/d/../"))
}

func simplifyPath(path string) string {
	pathSlice := strings.Split(path, "/")

	stack := make([]string, 0)

	for _, subPath := range pathSlice {
		switch subPath {
		case "", ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, subPath)
		}
	}

	var buf strings.Builder

	buf.WriteRune('/')
	buf.WriteString(strings.Join(stack, "/"))

	return buf.String()
}
