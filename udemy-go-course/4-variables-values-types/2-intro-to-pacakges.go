package main

import "fmt"

func main()  {
	n, err := fmt.Println("Hello world", 40, true)
	fmt.Println(n)
	fmt.Println(err)

	// Ignore return
	n2, _ := fmt.Println("Hello world", 50, true)
	fmt.Println(n2)
}