package main

import "fmt"

// each type has its own zero value
// let's see the zero values of the primitive types

var boolean bool        // zero value = false
var str string          // zero value = ""
var numberInt int       // zero value = 0
var numberFloat float64 // zero value = 0.0

func main() {
	fmt.Printf("%v, %T\n", boolean, boolean)
	fmt.Printf("%v, %T\n", str, str)
	fmt.Printf("%v, %T\n", numberInt, numberInt)
	fmt.Printf("%v, %T\n", numberFloat, numberFloat)

}
// zero values
// Made in Go playground
// https://go.dev/play/p/iStcbERXsD2