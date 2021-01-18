package main

import "fmt"

// Use var for:
// Zero value
// Package declaration
var y2 string
var z4 int

func main()  {
	fmt.Println("Printing the value of y", y2, "ending.")
	fmt.Printf("%T\n", y2)

	y2 = "Bond, James Bond"
	fmt.Println(y2)
	fmt.Printf("%T", y2)
	fmt.Println(z4)
	fmt.Printf("%T", z4)
}
