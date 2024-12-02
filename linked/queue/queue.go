package queue

import (
	"errors"
)

// QueueNode represents a node in the queue
type QueueNode struct {
	Data int
	Next *QueueNode
}

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
}

func InitQueue() *Queue {
	return &Queue{
		Head: nil,
		Tail: nil,
	}
}

func Enqueue(q *Queue, x int) {
	node := &QueueNode{
		Data: x,
		Next: nil,
	}
	if q.Head == nil {
		// Queue is empty
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

func Dequeue(q *Queue, value *int) error {
	if q.Head == nil {
		return errors.New("Queue is empty") // Queue is empty
	}
	*value = q.Head.Data
	q.Head = q.Head.Next
	if q.Head == nil {
		// If the queue becomes empty, set Tail to nil
		q.Tail = nil
	}
	return nil
}

func IsQEmpty(q *Queue) bool {
	return q.Head == nil
}

func Head(q *Queue) int {
	if q.Head == nil {
		return 0 // Queue is empty
	}
	return q.Head.Data
}
