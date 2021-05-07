// Singly Linked List:
// 1. add_at_start
// 2. using new to allocate memory for a node
// 3. creating a method receiver for the list

package main

import "fmt"

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	var head *IntList
	add_at_start(&head, 10)
	add_at_start(&head, 20)
	sum := head.Sum()
	fmt.Println(sum)
}

func add_at_start(head **IntList, val int) {
	// tmpNode := make(*IntList, len(*IntList)) // doesn't work
	tmpNode := new(IntList)
	tmpNode.Value = val
	tmpNode.Tail = *head
	*head = tmpNode
}
