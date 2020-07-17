package main

import "fmt"

func main() {

	var foo string
	var bar = "A string literal"
	gak := `A raw string literal, "can have all sorts for formatting",
	that can do anything
	
	cool!`

	fmt.Println("foo is:", foo)
	fmt.Println("bar is:", bar)
	fmt.Println("gar is:", gak)

}
