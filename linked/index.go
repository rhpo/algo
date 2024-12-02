package main

import (
	"fmt"
	. "linked/list"
)

var l PList
var p PList

func main() {

	n := 0
	for n <= 0 {
		fmt.Print("Enter n: ")
		fmt.Scanf("%d", &n)
	}

	x := 0
	for i := 0; i < n; i++ {
		fmt.Print("Enter value: ")
		fmt.Scanf("%d", &x)

		p = &Node{Data: x}
		p.Next = l
		l = p
	}

	var q PList

	q = l
	for q.Next != nil {
		q = q.Next
	}

	for i := 0; i < n; i++ {
		fmt.Print("Enter value: ")
		fmt.Scanf("%d", &x)

		q.Next = &Node{Data: x}
		q = q.Next
	}

	Display(l)
}
