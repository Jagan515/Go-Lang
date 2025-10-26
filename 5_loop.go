package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		//continue
		if i == 3 {
			continue
		}

		fmt.Println("Iteration number:", i)
	}

	// while loop in go
	j := 1
	for j <= 5 {
		//break statement
		if j == 4 {
			break
		}
		fmt.Println("While loop iteration:", j)
		j++
	}
	for i := range "Golang" {
		fmt.Println("Character index:", i)
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite Loop")
	// }
}
