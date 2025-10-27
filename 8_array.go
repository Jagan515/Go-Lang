package main

import "fmt"

func main() {
	var arr [5]int // declare an array of 5 integers
	arr[0] = 10    // assign value to the first element
	arr[1] = 20    // assign value to the second element
	arr[2] = 30    // assign value to the third element
	arr[3] = 40    // assign value to the fourth element
	arr[4] = 50    // assign value to the fifth element
	fmt.Println("Array elements:", arr)

	// declare and initialize an array in one line
	arr2 := [3]string{"Go", "Python", "Java"}
	fmt.Println("Array 2 elements:", arr2)

	// iterate over the array
	for i, v := range arr2 {
		fmt.Printf("Element at index %d: %s\n", i, v)
	}

	// get the length of the array
	fmt.Println("Length of arr:", len(arr))

	// multidimensional array
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D Array (Matrix):", matrix)

	// access elements in the multidimensional array
	fmt.Println("Element at (1,2):", matrix[1][2])
	// access all elements in the first row
	fmt.Println("First row elements:", matrix[0])

	//acess all elements
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at (%d,%d): %d\n", i, j, matrix[i][j])
		}
	}

	//fix size than array otherwise use slice
	// constant time access to elements
}
