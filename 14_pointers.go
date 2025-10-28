package main

import "fmt"

func changedFunction(x int) {
	x = 10
	fmt.Println("Value of x inside changedFunction:", x)
}
func pointerFunction(y *int) {
	*y = 30
	fmt.Println("Value of y inside pointerFunction:", *y)
}
func main() {

	a := 5
	fmt.Println("Value of a before changedFunction call:", a)
	changedFunction(a)
	fmt.Println("Value of a after changedFunction call:", a)

	//pointer example

	b := 20
	fmt.Println("Value of b before pointerFunction call:", b)
	pointerFunction(&b)
	fmt.Println("Value of b after pointerFunction call:", b)
}
