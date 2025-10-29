package main

import "fmt"

func printSlice[T any, s string](items []T, name s) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

func printsecond[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// Lifo
type Stack[T any] struct {
	element []T
}

func main() {
	names := []string{"GoLang", "Python", "C++"}
	nums := []int{1, 2, 3, 4, 5, 6}
	printSlice(names, "Jagan Pradhan")
	printSlice(nums, "Go Lang")

	mystack := Stack[int]{
		element: []int{1, 2, 3, 4, 5, 6},
	}

	mystack2 := Stack[string]{
		element: []string{"GoLang", "Python", "Cpp"},
	}

	fmt.Println(mystack)
	fmt.Println(mystack2)
}
