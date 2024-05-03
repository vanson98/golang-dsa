package stk

import (
	"fmt"
)

type Stack[T any] struct {
	data []T
}

func TestStack() {
	// genertic type must always supply all type parameter upon instantiation -- https://stackoverflow.com/questions/68166558/generic-structs-with-go
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Pop()
	fmt.Println(stack.data)
	topStack := stack.Top()
	fmt.Printf("\nTop: %v", topStack)

}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("Stack Empty 123")
	} else {
		topElement := s.data[len(s.data)-1]
		s.data = s.data[0 : len(s.data)-1]
		return topElement
	}

}

func (s *Stack[T]) Top() T {
	if s.Size() == 0 {
		panic("Stack Empty")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}
