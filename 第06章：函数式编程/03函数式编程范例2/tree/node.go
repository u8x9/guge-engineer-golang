package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
