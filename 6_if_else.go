package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// if-else statement
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}

	// if-else if-else ladder
	marks := 85
	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 80 {
		fmt.Println("Grade: B")
	} else if marks >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// || and && operators
	age := 25
	citizen := true
	if age >= 18 && citizen {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not eligible to vote")
	}

	if temp := 30; temp > 25 {
		fmt.Println("It's a hot day")
	} else {
		fmt.Println("It's a cool day")
	}
	// go does not support ternary operator
}
