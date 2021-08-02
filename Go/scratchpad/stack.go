package main

import (
	"fmt"
	"log"
)

type Stack interface {
	IsEmpty() bool
	Top() interface{}
	Push(interface{}, interface{})
	Pop()
	Print()
}

type StackInt struct {
	stack []int
}

func (s StackInt) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s StackInt) Top() interface{} {
	return s.stack[len(s.stack)-1]
}

func (s *StackInt) Push(element ...int) {
	s.stack = append(s.stack, element...)
}

func (s *StackInt) Pop() error {
	if s.IsEmpty() {
		return fmt.Errorf("stack is empty")
	} else {
		s.stack = s.stack[:len(s.stack)-1]
		return nil
	}
}

func (s StackInt) Print() {
	fmt.Println(s.stack)
}

func NewIntStack() StackInt {
	var s []int
	return StackInt{
		stack: s,
	}
}

func (s *StackInt) Erase(ntop int) {
	eraseCount := (len(s.stack) - ntop)
	var empty []int
	if eraseCount < 0 {
		log.Println("Erase count greater than length of stack")
		s.stack = empty
	} else {
		s.stack = s.stack[:eraseCount]
	}
}

func main() {
	intStack := NewIntStack()
	intStack.Push(10, 20, 30, 40)
	fmt.Println(intStack.Top())
	intStack.Erase(6)
	intStack.Print()
}
