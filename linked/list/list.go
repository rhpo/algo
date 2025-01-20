package list

import (
	"fmt"
	"strings"
)

func Read(x *int) {
	fmt.Scanf("%d", x)
}

func Write(args ...interface{}) {
	// No arguments: print a newline.
	if len(args) == 0 {
		fmt.Println()
		return
	}

	// If there are multiple arguments, join them with a space and print.
	var parts []string
	for _, arg := range args {
		parts = append(parts, fmt.Sprint(arg)) // Convert each argument to a string
	}
	fmt.Println(strings.Join(parts, " "))
}

type Node struct {
	Data int
	Next PList
}

type PList *Node

func Display(ll PList) {
	h := ll

	for h != nil {
		print(h.Data, " -> ")
		h = h.Next
	}

	println("NIL")

}

// AddToEnd adds a new node with the given data to the end of the list
// func (ll *LinkedList) AddToEnd(data int) {
// 	newNode := &Node{Data: data, Next: nil}
// 	if ll.Head == nil {
// 		ll.Head = newNode
// 	} else {
// 		current := ll.Head
// 		for current.Next != nil {
// 			current = current.Next
// 		}
// 		current.Next = newNode
// 	}
// }

// // AddToStart adds a new node with the given data to the start of the list
// func (ll *LinkedList) AddToStart(data int) {
// 	newNode := &Node{Data: data, Next: ll.Head}
// 	ll.Head = newNode
// }

// Remove deletes the first occurrence of the given data from the list
// func (ll *LinkedList) Remove(data int) error {
// 	if ll.Head == nil {
// 		return errors.New("list is empty")
// 	}

// 	if ll.Head.Data == data {
// 		ll.Head = ll.Head.Next
// 		return nil
// 	}

// 	current := ll.Head
// 	for current.Next != nil && current.Next.Data != data {
// 		current = current.Next
// 	}

// 	if current.Next == nil {
// 		return errors.New("data not found in the list")
// 	}

// 	current.Next = current.Next.Next
// 	return nil
// }

// Traverse returns a slice containing all elements in the list
// func (ll *LinkedList) Traverse() []int {
// 	result := []int{}
// 	current := ll.Head
// 	for current != nil {
// 		result = append(result, current.Data)
// 		current = current.Next
// 	}
// 	return result
// }
