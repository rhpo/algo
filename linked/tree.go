package main

import (
	. "linked/tree"
)

func LMinTree(tree *Node) int {

	if tree == nil {
		return 0
	} else {
		if LeftChild(tree) != nil {
			return 1 + LMinTree(LeftChild(tree))
		} else {
			return 0
		}
	}

}

func main() {

	var t *Tree = NewTree()

	Insert(t, 0)
	Insert(t, 7)
	Insert(t, 2)
	Insert(t, 1)

	DisplayTree(t)

	levelMin := LMinTree(t.Root)

	println(levelMin)
}
