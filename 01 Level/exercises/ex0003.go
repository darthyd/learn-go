package main

import "fmt"

// 1. declare 3 package-level scope variables (x, y, z) of type int, string and bool
// with the values 42, "James Bond" and true
// 2. on the main function use Sprintf to assigns these to a single variable named s
// with the shortcut operator
// 3. print the s value

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}

// exercise 0003
// Made on Go Playground
// https://go.dev/play/p/OYmLbWU-1TX