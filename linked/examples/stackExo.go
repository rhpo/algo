package main

import (
	"linked/queue"
	. "linked/stack"
	. "linked/utils"
)

func displayRev(s *Stack) {
	M := InitStack()

	var x int
	for !IsSEmpty(s) {
		Pop(&s, &x)
		Push(&M, x)
	}

	for !IsSEmpty(M) {
		Pop(&M, &x)
		Push(&s, x)
		Write(x, " ")
	}

	Write()
}

func occ(s *Stack, z int, n *int) {
	*n = 0
	M := InitStack()

	var x int
	for !IsSEmpty(s) {
		Pop(&s, &x)
		Push(&M, x)
	}

	for !IsSEmpty(M) {
		Pop(&M, &x)
		Push(&s, x)
		if x == int(z) {
			*n++
		}
		Write(x, " ")
	}

	Write()
}

func removeMul(s **Stack, z int) {

	M := queue.InitQueue()
	var x int

	for !IsSEmpty(*s) {
		Pop(s, &x)
		if Top(*s) != x {
			queue.Enqueue(M, x)
		}
	}

	xp := x

	for !queue.IsQEmpty(M) {
		queue.Dequeue(M, &xp)
		Write("Pushing", xp, "back to the main stack...\n")
		Push(s, xp)
	}

}

func sortAscending(s **Stack) {

	// sort the elements of the stack S in ascending order

	M := InitStack()
	var x, y int

	for !IsSEmpty(*s) {
		Pop(s, &x)
		for !IsSEmpty(M) && x < Top(M) {
			Pop(&M, &y)
			Push(s, y)
		}
		Push(&M, x)
	}

	for !IsSEmpty(M) {
		Pop(&M, &x)
		Push(s, x)
	}

}

func insertSorted(s *Stack, x int) {

	// insert element x in the sorted stack s

	M := InitStack()
	var y int

	for !IsSEmpty(s) && Top(s) > x {
		Pop(&s, &y)
		Push(&M, y)
	}

	Push(&s, x)

	for !IsSEmpty(M) {
		Pop(&M, &y)
		Push(&s, y)

	}

}

func InvertStack(s **Stack) {

	var x int
	R := InitStack()

	for !IsSEmpty(*s) {
		Pop(s, &x)
		Push(&R, x)
	}

	*s = R

}

func main() {

	var S *Stack = InitStack()

	for i := 0; i < 5; i++ {
		Push(&S, int(i*2))
	}

	DisplayStack(S)

	InvertStack(&S)

	DisplayStack(S)

	var x, res int
	Write("Enter number: ")
	Read(&x)

	occ(S, x, &res)

	Write("Occurences: ", res, "\n")

	removeMul(&S, x)
	DisplayStack(S)

	// Write("Reverse: ")
	// displayRev(S)

	// DisplayStack(S)

}
