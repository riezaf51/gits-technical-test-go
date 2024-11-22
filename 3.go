package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input string
	input, _ = reader.ReadString('\n')
	if isBalancedParentheses(input) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isBalancedParentheses(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r' {
			continue
		}

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) > 0 && s[i] == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
			continue
		}

		if len(stack) > 0 && s[i] == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
			continue
		}

		if len(stack) > 0 && s[i] == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
			continue
		}

		return false
	}

	return len(stack) == 0
}
