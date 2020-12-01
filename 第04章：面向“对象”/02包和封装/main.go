package main

import (
	"fmt"

	"github.com/u8x9/guge-engineer-golang/第04章：面向“对象”/02包和封装/tree"
)

// 结构体的创建
func createStruct() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
func main() {
	// createStruct()

	root := tree.NewNode(3)
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = &tree.Node{Value: 2}
	root.Traverse()

	fmt.Println()
}
