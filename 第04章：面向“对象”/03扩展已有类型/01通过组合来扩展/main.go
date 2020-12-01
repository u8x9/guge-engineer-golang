package main

import (
	"fmt"

	"github.com/u8x9/guge-engineer-golang/第04章：面向“对象”/02包和封装/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (n *myTreeNode) postOrder() {
	if n == nil || n.node == nil {
		return
	}
	left := myTreeNode{n.node.Left}
	left.postOrder()
	right := myTreeNode{n.node.Right}
	right.postOrder()
	n.node.Print()
}

func main() {
	root := tree.NewNode(3)
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = &tree.Node{Value: 2}

	root.Traverse()

	fmt.Println()

	myNodeRoot := &myTreeNode{root}
	myNodeRoot.postOrder()

	fmt.Println()
}
