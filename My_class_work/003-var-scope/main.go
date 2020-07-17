package main

import (
	"fmt"
)

var y = 43

//Declare the variable with the IDENTIFIER "z" of TYPE int
var z int

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	foo()
}

func foo() {
	z = 100
	fmt.Println("again:", y)
	fmt.Println("now z:", z)

}
