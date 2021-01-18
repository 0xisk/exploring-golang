package main

import "fmt"

var x = 43
var z int

func main()  {
	// BEST PRACTICE:
	// Limit the scope of your variables and as much as possible try to use shot-declaration operator
	var y = 42
	z = 32
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(z)
}
