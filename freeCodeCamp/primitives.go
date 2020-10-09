package main

import "fmt"

func main()  {
	// Bool declaration
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	// Logical test
	m := 1 == 1
	y := 1 == 2
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", y, y)

	// Initialization value for Bool
	var x bool
	fmt.Printf("%v, %T\n", x, x)
}
