package main

import (
	"fmt"
	"nqueens/nqueens"
)

func main() {
	for i := 1; ; i++ {
		sol := nqueens.HowManySolutionsConcurrent(i)
		fmt.Printf("Size %d: %d solutions.\n", i, sol)
	}

}
