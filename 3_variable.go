package main

import "fmt"

// mt is a standard library package used for formatted I/O operations, similar to printf and scanf in C.

func main() {
	var a int = 10
	var b string = "Golang"
	var c float32 = 5.67
	// Auto detect data type
	var d = 10.01

	// shorthand declaration
	e := "ShortHand Declaration"

	// If i want to declared a variable without assigning any value its compulsory to use var keyword
	var f int

	fmt.Println(f) // default value will be 0 for int
	f = 20
	fmt.Println(f)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
