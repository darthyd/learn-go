package main

import "fmt"

type hotdog int

var a hotdog = 11

func main() {
	fmt.Println("custom types x primitive types")
	b := 22
	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n\n", b, b)

	// b = a its not possible in this case cause
	// the subjacent type of a variable be the same
	// that the b type doesn't matter
	// i have to convert it to can this assign
	fmt.Println("converting types")
	b = int(a)
	fmt.Printf("%v, %t", b, b)

}


// DOCS: https://go.dev/tour/basics/13

// custom types and conversion of types
// Made on Go Playground
// https://go.dev/play/p/_T9MOLy5Y4F