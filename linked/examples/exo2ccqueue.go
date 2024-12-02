package main

import (
	"fmt"
	. "linked/queue"
	. "linked/stack"
	. "linked/utils"
)

func DisplayQueue(q *Queue) {
	var x int
	var q2 *Queue
	q2 = InitQueue()

	for !IsQEmpty(q) {
		Dequeue(q, &x)
		fmt.Print(x, " ")
		Enqueue(q2, x)
	}

	Write()

	for !IsQEmpty(q2) {
		Dequeue(q2, &x)
		Enqueue(q, x)
	}

}

func Decompose(q *Queue, n int) {

	var s *Stack = InitStack()

	var digit int

	for n != 0 {
		digit = n % 10
		Push(&s, digit)
		n /= 10
	}

	for !IsSEmpty(s) {
		Pop(&s, &digit)
		Enqueue(q, digit)
	}

}

func Compose(q *Queue) int {
	result := 0
	digit := 0

	for !IsQEmpty(q) {
		Dequeue(q, &digit)
		result = 10*result + digit
	}

	return result
}

func Add(q1 *Queue, q2 *Queue, out *Queue) {

	Decompose(out, Compose(q1)+Compose(q2))

}

func Complete(f FQueue) {

	fp := InitQueue()
	Fqueuenode x;

	for not isQEMPTY(f) {
		dequeue(f, x)
		Fqueuenode.info.rep = decompose(Fqueuenode.info.value)
		enqueue(fp, Fqueuenode)
	}

	for not isQeumpty(fp) {
		dequeue(fp, x)
		enqueue(f, x)
	}

}

func main() {

	var t FQueue

	t = &FQueueNode{
		Info: TypeElt{
			value: 123,
			rep: nil,
		},
		Next: &FQueueNode{
			Info: TypeElt{
				value: 456,
				rep: nil,

			},
			Next: nil,
		},
	}

	q1 := InitQueue()
	q2 := InitQueue()
	qr := InitQueue()

	n1 := 0
	n2 := 0

	Write("Enter N1 ")
	Read(&n1)

	Write("Enter N2 ")
	Read(&n2)

	Decompose(q1, n1)
	Decompose(q2, n2)

	// fmt.Println("Q1:")
	// DisplayQueue(q1)

	// fmt.Println("Q2:")
	// DisplayQueue(q2)

	Add(q1, q2, qr)

	fmt.Println("Result:")
	DisplayQueue(qr)

}
