package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func (n *Node) InsertAtStart(data int) {
	newNode := createNewNode(data)
	newNode.next = n
	n = newNode
}

func (n *Node) InsertAtEnd(data int) {

}

func (n *Node) InsertAtPos(data, pos int) {

}

func createNewNode(data int) *Node {
	return &Node{data: data}
}

func main() {
	var head *Node
	node.InsertAtStart(10)
	fmt.Println(head)
}
