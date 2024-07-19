package main

import "fmt"

func main() {
	//simple Values

	// int
	fmt.Println(1 + 1)

	// string
	fmt.Println("Hello, World!")

	// bool
	fmt.Println(true)
	fmt.Println(false)

	// float
	fmt.Println(3.14)

	// Arthmetic Operators
	fmt.Println(1 + 1)
	fmt.Println(1 - 1)
	fmt.Println(1 * 1)
	fmt.Println(1 / 1)
	fmt.Println(1 % 1)

	// Comparison Operators
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)
	fmt.Println(1 > 1)
	fmt.Println(1 < 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 <= 1)

	// Logical Operators
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Bitwise Operators
	fmt.Println(1 & 1)
	fmt.Println(1 | 1)
	fmt.Println(1 ^ 1)
	fmt.Println(1 << 1)
	fmt.Println(1 >> 1)

	// Assignment Operators
	a := 1
	a += 1
	fmt.Println(a)

	// Increment and Decrement Operators
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// Ternary Operator
	a = 1
	if a == 1 {
		fmt.Println("a is 1")
	} else {
		fmt.Println("a is not 1")
	}

}