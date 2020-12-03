package main

import (
	"fmt"

	"github.com/u8x9/guge-engineer-golang/第10章：channel/03使用Channel进行树的遍历/tree"
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

	c := root.TraverseWithChannel()
	maxNode := 0
	for n := range c {
		if n.Value > maxNode {
			maxNode = n.Value
		}
	}
	fmt.Println("Max node:", maxNode)
}
