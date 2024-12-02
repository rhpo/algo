package main

import (
	. "linked/stack"
	"math/rand/v2"
)

func SortAscending(s **Stack) {

	// sort the elements of the stack S in ascending order

	var x, y int
	M := InitStack()

	for !IsSEmpty(*s) {
		Pop(s, &x)

		for !IsSEmpty(M) && Top(M) > x {
			Pop(&M, &y)
			Push(s, y)
		}

		Push(&M, x)
	}

	*s = M

}

// سكران طايح فالدروج
// Let A be an array of n integers (n≤100). Write an algorithm to put the even numbers at the beginning
// of the array, then the odd numbers while preserving the order of appearance of the even numbers
// within the initial array, and reverse the order of the odd numbers.
// Example : A : 2 5 4 6 7 1 9 3 12 15 17 0 After A : 2 4 6 12 0 17 15 3 9 1 7 5
func SortEvenOdd(A []int) {

	// sort the elements of the array A in the following way:
	// - put the even numbers at the beginning of the array
	// - put the odd numbers at the end of the array in reverse order

	var i, j, n int
	var temp int

	n = len(A)

	for i = 0; i < n; i++ {
		if A[i]%2 == 0 {
			j = i
			for j > 0 && A[j-1]%2 != 0 {
				temp = A[j]
				A[j] = A[j-1]
				A[j-1] = temp
				j--
			}
		}
	}

	for i = 0; i < n; i++ {
		if A[i]%2 != 0 {
			j = i
			for j < n-1 && A[j+1]%2 != 0 {
				temp = A[j]
				A[j] = A[j+1]
				A[j+1] = temp
				j++
			}
		}
	}

}

func main() {

	var S *Stack = InitStack()

	value := rand.IntN(100)
	for i := 0; i < 5; i++ {
		Push(&S, int(value))
		value = rand.IntN(100)
	}

	DisplayStack(S)

	SortAscending(&S)

	DisplayStack(S)

}
