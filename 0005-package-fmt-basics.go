package main

import "fmt"

func main() {
	string := "formated string"

	fmt.Print("Print: just print something")

	fmt.Println("Println: print something and go to a new line")

	fmt.Printf("Printf: print a %v", string)

	// Prints the content of any writter interface, a file, a server response...
	fmt.Fprint("an io.Writter type")

	// Good to know the Print, Println and Printf fucntion returns two values when use
	// the number of bytes that was evaluated and for second any errors

	bytes, error := fmt.Println("any")
	fmt.Println(bytes, error)

	// differente of the other Sprint dont give a space between the given strings
	// instead all values is not a string
	string = fmt.Sprint("a function that returns the given", "string")
	fmt.Println(string)

	// we could use the variants Sprintln, Sprintf, Fprintln and Fprintf too

}


//https://pkg.go.dev/fmt