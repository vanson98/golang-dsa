package stk

import (
	"fmt"
	"regexp"
	"strings"
)

func TestInFixToPostfix() {
	infixToPostFix("A*B+C*((D-E)+F)/G + 234")
}

func formatExpression(expression string) []string {
	strings.Replace(expression, " ", "", -1)

	//https://golangbyexample.com/regex-replace-string-golang/
	re := regexp.MustCompile(`\+|\-|\*|\/|\%|\^|\)|\(`)
	expression = re.ReplaceAllString(expression, " $0 ")
	expression = strings.Replace(expression, "  ", " ", -1)
	expression = strings.Trim(expression, " ")
	return strings.Split(expression, " ")
}

func getPriority(operator string) int {
	if operator == "*" || operator == "/" || operator == "%" {
		return 2
	} else if operator == "+" || operator == "-" {
		return 1
	}
	return 0
}

func isOperator(element string) bool {
	re := regexp.MustCompile(`\+|\-|\*|\/|\%|\^||\(`)
	return re.MatchString(element)
}

func isOperand(element string) bool {
	re := regexp.MustCompile(`^\d+$|^([a-z]|[A-Z])$`)
	return re.MatchString(element)
}

func infixToPostFix(infix string) {
	elements := formatExpression(infix)
	stack := Stack[string]{}
	var stringBuilder strings.Builder
	// if for loop with infix -> each elemenet is rune UTF-32 int32 (byte-rune)   https://www.ekwbtblog.com/entry/2023/03/20/080011

	for _, v := range elements {
		if isOperand(v) {
			stringBuilder.WriteString(v)
		} else if isOperator(v) {
			if v == ")" {
				for {
					elm := stack.Pop()
					if elm == "(" {
						break
					}
					stringBuilder.WriteString(v)
				}

			} else {
				top := stack.Top()
				topElementPriority := getPriority(top)
				currentOperatorPriority := getPriority(v)
				if topElementPriority >= currentOperatorPriority {
					stringBuilder.WriteString(stack.Pop())
				}
				stack.Push(v)
			}
		}
	}
	for !stack.IsEmpty() {
		opert := stack.Pop()
		stringBuilder.WriteString(opert)
	}
	fmt.Println(stringBuilder.String())
}
