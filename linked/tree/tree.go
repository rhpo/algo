package tree

import (
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}

func NewTree() *Tree {
	return &Tree{}
}

func Height(tree *Tree) int {
	return height(tree.Root)
}

func height(node *Node) int {
	if node == nil {
		return 0
	}

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func ClassicPrintTree(tree *Tree) {
	if tree.Root == nil {
		fmt.Println("Tree is empty")
		return
	}

	printNode(tree.Root)
}

func LeftChild(tree *Node) *Node {
	return tree.Left
}

func RightChild(tree *Node) *Node {
	return tree.Right
}

func printNode(node *Node) {
	if node == nil {
		return
	}

	printNode(node.Left)
	fmt.Println(node.Data)
	printNode(node.Right)
}

func Insert(tree *Tree, data int) {
	if tree.Root == nil {
		tree.Root = NewNode(data)
	} else {
		insertNode(tree.Root, data)
	}
}

func insertNode(node *Node, data int) {
	if data < node.Data {
		if node.Left == nil {
			node.Left = NewNode(data)
		} else {
			insertNode(node.Left, data)
		}
	} else {
		if node.Right == nil {
			node.Right = NewNode(data)
		} else {
			insertNode(node.Right, data)
		}
	}
}

func DisplayTree(tree *Tree) {
	showTree(tree)
}
