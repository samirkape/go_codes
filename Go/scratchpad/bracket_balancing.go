package main

import "fmt"

type Brackets struct {
	bracket []byte
}

func (b *Brackets) Push(ch byte) {
	b.bracket = append(b.bracket, ch)
}

func (b Brackets) Pop() byte {
	return b.bracket[len(b.bracket)-1]
}

func isBalancedF(input string, stack Brackets) bool {
	for i := 0; i < len(input); i++ {
		e := input[i]
		if e == '{' || e == '(' || e == '[' {
			stack.Push(e)
		} else if e == '}' {
			b := stack.Pop()
			if b != '{' {
				return false
			}
		} else if e == ')' {
			b := stack.Pop()
			if b != '(' {
				return false
			}
		} else if e == ']' {
			b := stack.Pop()
			if b != '[' {
				return false
			}
		}
	}
	return true
}

func (stack Brackets) isBalanced(input string) bool {
	for i := 0; i < len(input); i++ {
		e := input[i]
		if e == '{' || e == '(' || e == '[' {
			stack.Push(e)
		} else if e == '}' {
			b := stack.Pop()
			if b != '{' {
				return false
			}
		} else if e == ')' {
			b := stack.Pop()
			if b != '(' {
				return false
			}
		} else if e == ']' {
			b := stack.Pop()
			if b != '[' {
				return false
			}
		}
	}
	return true
}

func isValidClose(stack Brackets, bracket byte) bool {
	b := stack.Pop()
	return b == bracket
}

func (stack Brackets) isBalancedSw(input string) bool {
	for i := 0; i < len(input); i++ {
		e := input[i]
		switch e {
		case '{', '(', '[':
			stack.Push(e)
		case '}':
			if isValidClose(stack, '{') == false {
				return false
			}
		case ')':
			if isValidClose(stack, '(') == false {
				return false
			}
		case ']':
			if isValidClose(stack, '[') == false {
				return false
			}
		}
	}
	return true
}

func main() {
	input := "()()[[]]"
	var buffer []byte
	stack := Brackets{
		bracket: buffer,
	}
	result := stack.isBalancedSw(input)
	fmt.Println(result)
}
