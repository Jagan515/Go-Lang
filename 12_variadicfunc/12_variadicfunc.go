package main

import "fmt"

// variadic function example

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//PrintLn is a variadic function that prints all arguments eg fmt.Println("a", "b", "c")

// anytype variadic function example passing any type of arguments

func anyTypePrint(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	result1 := sum(1, 2, 3)
	fmt.Println("Sum of 1, 2, 3 is:", result1)

	result2 := sum(10, 20, 30, 40, 50)
	fmt.Println("Sum of 10, 20, 30, 40, 50 is:", result2)

	nums := []int{5, 10, 15}
	result3 := sum(nums...)
	fmt.Println("Sum of slice nums is:", result3)

	//string and int and float64 types passed to anyTypePrint

	anyTypePrint("Hello", 42, 3.14, true)

	// we want to pass slice to anyTypePrint
	mixedSlice := []interface{}{"Golang", 100, 9.81, false}
	anyTypePrint(mixedSlice...)
}
