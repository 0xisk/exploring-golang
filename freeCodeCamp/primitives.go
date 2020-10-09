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

	// Integers
	i1 := 42
	fmt.Printf("%v, %T\n", i1, i1)

	var i2 uint16 = 42
	fmt.Printf("%v, %T\n", i2, i2)

	// Invalid types error
	/*var i3 int = 10
	var i4 int8 = 3
	fmt.Println(i3 + i4)*/

	// Operators
	i5 := 10  // 1010
	i6 := 3   // 0011
	fmt.Println(i5 & i6) // 0010 ~ 2
	fmt.Println(i5 | i6) // 1011 ~ 11
	fmt.Println(i5 ^ i6) // 1001 ~ 9
	fmt.Println(i5 &^ i6) // 0100 ~ 8

	// Bit shifting
	i7 := 8 // 2^3
	fmt.Println(i7 << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(i7 >> 3) // 2^3 / 2^3 = 2^0 = 1
}
