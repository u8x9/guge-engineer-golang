package main

import (
	"fmt"

	"github.com/u8x9/guge-engineer-golang/第06章：函数式编程/03函数式编程范例2/tree"
)

func main() {
	root := &tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.NewNode(2)

	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
