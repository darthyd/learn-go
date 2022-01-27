package main

import "fmt"

// package level scope declaration
// you have to use the var keyword
var pckLevelScope = "package level scope variable"

// you can declare a variable without initialize it with a value
// but the attribution can only be made inside a block scope
// in this case you have to give this a type
var nonInit string

func main() {

	// initialize and declare the var value
	nonInit = "initialize and declare"

	fmt.Println(nonInit, "\n") // you can go to a new line with a \n into a string

	// vars declaration
	// in var declaration, the type is optional
	// this variables have block scope
	// you dont need to use the var keyword
	string := "string" // type is string
	number := 16       // type is int
	bool := true       // type is bool

	fmt.Println(string, number, bool)

	// vars attribution
	// in var attribution, the type is mandatory
	// this means you cannot change the type of decalred variable
	// you cannot use := to assign a value to a variable
	// that has already been declared
	string = "new string"

	// declaration and attribution in one line
	string, string2 := "string", "another string"

	// checking types
	fmt.Println("\nchecking types:")
	fmt.Printf("string2: %v, %T\n", string2, string2)
	fmt.Printf("number: %v, %T\n", number, number)
	fmt.Printf("bool: %v, %T\n", bool, bool)

}

// RESERVED WORDS (KEYWORDS)

// var, func, package, import, return,
// break, continue, for, if, else, switch,
// case, default, goto, fallthrough, range,
// type, interface, map, struct, chan, select,
// else, defer, go, range

// var declarations, attributions, scope and primitive types
// Made in Go playground
// https://go.dev/play/p/XR-HBZ501l3