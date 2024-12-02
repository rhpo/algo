package stack

import (
	. "linked/utils"
)

type Stack struct {
	Data int
	Next *Stack

	Tail *Stack
}

func InitStack() *Stack {
	return nil
}

func Push(s **Stack, x int) {

	node := &Stack{
		Data: x,
		Next: nil,
	}

	if IsSEmpty(*s) {
		*s = node
		(*s).Tail = node
	} else {
		P := s

		for (*P).Next != nil {
			P = &(*P).Next
		}

		(*P).Next = node
		(*s).Tail = node

	}

}

func DisplayStack(s *Stack) {
	if IsSEmpty(s) {
		Write("Stack is empty")
		return
	}

	Write("Stack elements: ")
	for !IsSEmpty(s) {
		Write(s.Data, " ")
		s = s.Next
	}

	Write()
}

// Pop removes the top element from the stack
func Pop(s **Stack, value *int) {
	// removes last element of the stack
	if IsSEmpty(*s) {
		return
	}

	if (*s).Next == nil {

		*value = (*s).Data
		*s = nil
		return
	}

	P := s

	for (*P).Next.Next != nil {

		P = &(*P).Next

	}

	*value = (*P).Next.Data
	(*P).Next = nil
	(*s).Tail = *P

}

// IsSEmpty checks if the stack is empty
func IsSEmpty(s *Stack) bool {
	return s == nil
}

// Top retrieves the top element of the stack
func Top(s *Stack) int {
	if s == nil {
		return 0 // Stack is empty
	}
	return (*s).Tail.Data
}
