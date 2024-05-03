package stk

import (
	"fmt"
)

func CheckParenthesis(input string) {
	stack := Stack[string]{}
	for _, v := range input {
		bracket := string(v)
		if bracket == "(" || bracket == "{" || bracket == "[" {
			stack.Push(bracket)
			continue
		}
		if stack.IsEmpty() {
			fmt.Print("Not Match")
			return
		}
		if bracket == ")" && stack.Top() == "(" {
			stack.Pop()
		} else if bracket == "}" && stack.Top() == "{" {
			stack.Pop()
		} else if bracket == "]" && stack.Top() == "[" {
			stack.Pop()
		}
	}
	if stack.IsEmpty() {
		fmt.Print("Match")
	} else {
		fmt.Println("Not Match")
	}
}
