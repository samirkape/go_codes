package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	data int
	next *Node
}

type Position func(int, **Node)

func InsertAt(pos Position, node **Node, data ...int) {
	for _, e := range data {
		pos(e, node)
	}
}

func beginning(data int, node **Node) {
	newNode := createNewNode(data)
	if reflect.ValueOf(*node).IsZero() {
		*node = newNode
		return
	}
	newNode.next = (*node)
	*node = newNode
}

func end(data int, node **Node) {
	tmpNode := *node
	newNode := createNewNode(data)
	if reflect.ValueOf(*node).IsZero() {
		*node = newNode
	} else {
		for tmpNode.next != nil {
			tmpNode = tmpNode.next
		}
		tmpNode.next = newNode
	}
}

func InsertAtPosition(data, position int, node **Node) {
	if position == 1 {
		beginning(data, node)
	}
	tmpNode := *node
	newNode := createNewNode(data)
	for pc := 1; pc < position-1; pc++ {
		tmpNode = tmpNode.next
	}
	newNode.next = tmpNode.next
	tmpNode.next = newNode
}

func createNewNode(data int) *Node {
	return &Node{data: data}
}

func PrintList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.data)
		head = head.next
	}
	fmt.Println()
}

func main() {
	var head *Node
	blist := []int{30, 20, 10}
	InsertAt(beginning, &head, blist...)
	InsertAt(end, &head, 50)
	InsertAt(end, &head, 60, 70, 80, 90)
	InsertAtPosition(40, 4, &head)
	PrintList(head)
}
