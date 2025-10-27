package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello, Slices!")

	// declare and initialize a slice
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice elements:", slice)

	// append elements to the slice
	slice = append(slice, 60, 70)
	fmt.Println("Slice after appending elements:", slice)

	// iterate over the slice
	for i, v := range slice {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}

	// get the length of the slice
	fmt.Println("Length of slice:", len(slice))

	// create a slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := arr[1:4] // elements from index 1 to 3
	fmt.Println("Slice from array:", sliceFromArray)

	// modify the slice and see the effect on the underlying array
	sliceFromArray[0] = 20
	fmt.Println("Modified slice from array:", sliceFromArray)
	fmt.Println("Underlying array after modifying slice:", arr)

	// create a slice using make
	madeSlice := make([]string, 3) // slice of strings with length 3
	madeSlice[0] = "Go"
	madeSlice[1] = "Python"
	madeSlice[2] = "Java"
	fmt.Println("Slice created with make:", madeSlice)

	// demonstrate capacity
	capSlice := make([]int, 2, 5) // length 2, capacity 5
	fmt.Println("Initial capSlice:", capSlice, "Length:", len(capSlice), "Capacity:", cap(capSlice))
	capSlice = append(capSlice, 10, 20, 30)
	fmt.Println("capSlice after appending elements:", capSlice, "Length:", len(capSlice), "Capacity:", cap(capSlice))

	// copying slices
	originalSlice := []int{1, 2, 3}
	copiedSlice := make([]int, len(originalSlice))
	copy(copiedSlice, originalSlice)
	fmt.Println("Original slice:", originalSlice)
	fmt.Println("Copied slice:", copiedSlice)

	// slicing a slice
	subSlice := slice[2:5] // elements from index 2 to 4
	fmt.Println("Sub-slice:", subSlice)
	fmt.Println(subSlice[:2]) // first two elements of sub-slice
	fmt.Println(subSlice[1:]) // elements from index 1 to end of sub-slice
	// full slice expression
	fullSlice := slice[1:4:5] // from index 1 to 3 with capacity up to index 5
	fmt.Println("Full slice expression:", fullSlice)

	// multidimensional slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Slice (Matrix):", matrix)

	// access elements in the multidimensional slice
	fmt.Println("Element at (1,2):", matrix[1][2])
	// access all elements in the first row
	fmt.Println("First row elements:", matrix[0])

	// access all elements
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at (%d,%d): %d\n", i, j, matrix[i][j])
		}
	}
	// equality of slices
	sliceA := []int{1, 2, 3}
	sliceB := []int{1, 2, 3}
	// slices cannot be compared directly using '=='

	fmt.Println("Are sliceA and sliceB equal?", slices.Equal(sliceA, sliceB))
	// slices are more flexible than arrays

}
