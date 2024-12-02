package utils

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
