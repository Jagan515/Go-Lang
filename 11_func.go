package main

import "fmt"

// int is the return type of the function
func additionalFunction(a int, b int) int {
	// This is an additional function added to demonstrate changes
	return a + b
}

// multiple return values from a function

func getLanguage() (string, string, int, string, bool) {
	return "Go", "typescript", 2, "python", true
}

// function inside a function

func processvalue(fun func(a int) int) int {

	x := fun(69)
	return x

}

// returning a function from a function

func returnedFunction() func(a int) int {
	return func(a int) int {
		return a * a
	}
}

func main() {
	result := additionalFunction(5, 10)
	fmt.Println("The result of additionalFunction is:", result)

	// _ to ignore values as go lang does not support unused variables

	lang1, lang2, version, lang3, _ := getLanguage()

	// as we used _ to ignore the last return value, we won't have an unused variable error

	fmt.Println("Languages and version:", lang1, lang2, version, lang3)

	// create an anonymous function and pass it to processvalue

	process := func(a int) int {
		fmt.Println("Processing value:", a)
		return a * 2
	}
	x := processvalue(process)
	fmt.Println("Value of the anonymous funct", x)

	// get a function from a function and use it
	squareFunc := returnedFunction()
	squaredValue := squareFunc(7)
	fmt.Println("Squared value:", squaredValue)

}
