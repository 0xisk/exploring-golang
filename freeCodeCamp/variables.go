package main

import (
	"fmt"
)
// Variables Method 3
var i int = 42

// Variables Method 4
/*var (
	actorName string = "Iskander"
	companion string = "Andrew"
	doctorNumber int = 3
	season int = 11
)

var (
	counter int = 0
)
*/

// Global variable
var I int = 23

func main()  {

	// Shadowing
	fmt.Println(i)

	// Variables Method 1
	// var i int
	// i = 42

	// Variables Method 2
	// var x int = 50
	// x := 50

	var theURL string = "http://google.com"

	var i int = 50

	fmt.Println(i)
	fmt.Println(theURL)
	// fmt.Printf("%v,3 %T", i, i)
}
