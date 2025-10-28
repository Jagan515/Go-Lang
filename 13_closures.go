package main

import "fmt"

func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// we call a function go into the call stack and after execution it comes out of the call stack
// but in closure the inner function retains the state of the outer function even after the outer function has finished execution

func main() {
	increment := incrementer()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	increment()
	fmt.Println(increment())

	// create a new incrementer
	newIncrement := incrementer()
	fmt.Println(newIncrement())
	fmt.Println(newIncrement())
}
