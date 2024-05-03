package stk

import (
	"fmt"
	"strings"
)

func ReverseString(input string) {
	stack := Stack[rune]{}
	for _, v := range input {
		stack.Push(v)
	}
	bytes := make([]rune, len(input))
	for !stack.IsEmpty() {
		bytes = append(bytes, stack.Pop())
	}
	fmt.Print(string(bytes))
}

func ReverStringUsingStringBuilder(input string) {
	var stringBuilder strings.Builder

	stack := Stack[rune]{}
	for _, v := range input {
		stack.Push(v)
	}

	for !stack.IsEmpty() {
		stringBuilder.WriteString(string(stack.Pop()))
	}
	fmt.Print(stringBuilder.String())
}
