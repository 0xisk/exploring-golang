package main

import "fmt"

var y = 7
var z2 = "Shaken, not stirred"
var z3 string = `James said: "Shaken, not stirred".`

func main()  {
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z2)
	fmt.Printf(z3)
}