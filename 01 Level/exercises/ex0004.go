package main

import "fmt"

// 1. create a type with subjacent type int
// 2. create a variable for this type with the identifier x
// 3. print the value of x
// 4. print the type of x
// 5. assigns the value 42 to x
// 6. print the value of x

type customType int
var x customType

func main() {
	fmt.Printf("%v, %T \n", x, x)
	x = 42
	fmt.Print(x)
}

// exercise 0004
// Made on Go Playground
// https://go.dev/play/p/ObuXY4h15wI