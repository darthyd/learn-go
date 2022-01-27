package main

import "fmt"

func main() {

	// interpreted string literals
	fmt.Println("interpreted: Lucas\nJosé \"de\"\tOliveira\n\n")

	// raw string literals
	fmt.Println(`raw: "Lucas
	
	
	José \n \t de 
	
	
	Oliveira"`)

}


// interpreted and raw string literals
// Made on Go playground
// https://go.dev/play/p/OTUb74M3B-e
