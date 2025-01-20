package main

import (
	"fmt"
	. "linked/list"
)

var H PList

func LongestSub(H PList, Sx *int, Px *PList) {
	if H == nil {
		*Sx = 0
		*Px = nil
	} else {
		p := H

		maxSequenceLength := 0
		currentSequenceLength := 1
		var pCurrentMax PList = nil
		var pCurrentStart PList = nil

		for p.Next != nil {
			if pCurrentStart == nil {
				pCurrentStart = p
			}

			if p.Next.Data > p.Data {
				currentSequenceLength++
			} else {
				if currentSequenceLength > maxSequenceLength {
					maxSequenceLength = currentSequenceLength
					pCurrentMax = pCurrentStart
				}
				currentSequenceLength = 1
				pCurrentStart = nil
			}

			p = p.Next
		}

		// Check the last sequence
		if currentSequenceLength > maxSequenceLength {
			maxSequenceLength = currentSequenceLength
			pCurrentMax = pCurrentStart
		}

		*Px = pCurrentMax
		*Sx = maxSequenceLength
	}
}

func DetachLongest(H PList, Px PList) {

	h := H
	prev := h

	for h != nil && h != Px {
		prev = h
		h = h.Next
	}

	endList := h
	if h != nil {
		for endList != nil && endList.Next != nil && endList.Data < endList.Next.Data {
			endList = endList.Next
		}

		fmt.Printf("P1: %d, P2: %d\n", prev.Data, endList.Data)
		prev.Next = endList.Next
	}

}

func main() {

	var n int

	for n <= 0 {
		fmt.Printf("Enter n: ")
		fmt.Scanf("%d", &n)
	}

	for i := 0; i < n; i++ {

		value := 0

		p := &Node{}
		p.Next = nil

		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scanf("%d", &value)
		p.Data = value

		q := H

		if H == nil {
			H = p
		} else {
			for q.Next != nil {
				q = q.Next
			}

			q.Next = p
		}

	}

	Display(H)

	var max int
	var PM PList

	LongestSub(H, &max, &PM)

	Display(PM)

	fmt.Println("Before :")
	Display(H)

	DetachLongest(H, PM)

	fmt.Println("After :")
	Display(H)

}
