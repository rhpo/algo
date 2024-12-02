package main

import (
	"fmt"
	. "linked/list"
)

func main() {
	var H, P PList
	var n, i, data int

	for n <= 0 {
		print("Enter N=")
		fmt.Scanf("%d", &n)
	}

	for i = 0; i < n; i++ {
		print("Enter Value ", i, " :")
		fmt.Scanf("%d", &data)

		P = &Node{Data: data}
		P.Next = H

		H = P
	}

	Display(H)

	sum := 0
	Q := H

	for Q != nil {
		sum += Q.Data
		Q = Q.Next
	}

	println("The sum is", sum)

}
