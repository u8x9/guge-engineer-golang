package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}
