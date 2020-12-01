package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

func newTreeNode(value int) *treeNode {
	return &treeNode{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (tn *treeNode) print() {
	fmt.Print(tn.value, " ")
}
func (tn *treeNode) traverse() {
	if tn == nil {
		return
	}
	tn.left.traverse()
	tn.print()
	tn.right.traverse()
}

// 结构体的创建
func createStruct() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
func main() {
	// createStruct()

	root := newTreeNode(3)
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = &treeNode{value: 2}
	root.traverse()

	fmt.Println()
}
