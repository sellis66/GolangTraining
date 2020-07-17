package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("Outer loop i is now %v\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tInner Loop y is now: %v\n", j)

		}
	}
}

