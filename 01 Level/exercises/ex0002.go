package main

import "fmt"

// 1. declatre 3 package-level scope variables (x, y, z) of type int, string and bool
// 2. show the values that the compiler automatic have assigned to the variables

// 1.
var x int
var y string
var z bool

func main() {
	// 2.
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

// exercise 0002
// Made on Go Playground
// https://go.dev/play/p/A9HimCKv3dv