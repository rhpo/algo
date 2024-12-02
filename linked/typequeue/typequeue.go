package typequeue

import "errors"

type TypeElt struct {
	value int
	rep   *FQueue
}

type FQueueNode struct {
	Info TypeElt
	Next *FQueueNode
}

type FQueue struct {
	Head *FQueueNode
	Tail *FQueueNode
}

func InitQueue() *FQueue {
	return nil
}

func Enqueue(q *FQueue, x TypeElt) {
	node := &FQueueNode{
		Info: x,
		Next: nil,
	}

	if q == nil {
		// Queue is empty
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

func Dequeue(q *FQueue, value *TypeElt) error {
	if q.Head == nil {
		return errors.New("Queue is empty") // Queue is empty
	}
	*value = q.Head.Info
	q.Head = q.Head.Next
	if q.Head == nil {
		// If the queue becomes empty, set Tail to nil
		q.Tail = nil
	}
	return nil
}

func IsQEmpty(q *FQueue) bool {
	return q.Head == nil
}

func Head(q *FQueue) TypeElt {
	if q.Head == nil {
		return TypeElt{} // Queue is empty
	}
	return q.Head.Info
}
