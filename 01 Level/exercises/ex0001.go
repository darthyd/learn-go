package main

import "fmt"

// 1. use the shortcut operator to assigns the values 42, "James Bond" and true
// to 3 variables called x, y, z
// 2. print every variable in a different print statement
// 3. print every variable in the same print statement

func main() {
	// 1.
	x, y, z := 42, "James Bond", true

	// 2.
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// 3.
	fmt.Println(x, y, z)

}

// exercise 0001
// Made on Go Playground
// https://go.dev/play/p/rh-qloboXQF