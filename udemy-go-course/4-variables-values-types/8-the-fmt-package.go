package main

import "fmt"

var t = 42

func main()  {
	fmt.Println(t)
	fmt.Printf("%T\n", t)
	fmt.Printf("%b\n", t)
	fmt.Printf("%x\n", t)
	fmt.Printf("%#x\n", t)

	s := fmt.Sprintf("%#x\n", t)
	fmt.Println(s)
}
