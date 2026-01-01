package main

import (
	"fmt"
	"math"
)

const A string = "Hello"
const PI = 3.14

func main() {
	fmt.Println(A)

	const B = 400000000
	const C = 4e20 / B
	fmt.Println(C)

	fmt.Println(int64(C))

	fmt.Println(math.Sin(B))

	fmt.Println(PI)

	const D int = 1
	fmt.Println(D)

	const E = 3
	fmt.Println(E)

	const (
		F int = 6
		G     = 5.26
		I     = "Hello Go!"
	)

	fmt.Println(F, G, I)
}
