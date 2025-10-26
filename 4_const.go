package main

import "fmt"

func main() {
	const pi float32 = 3.14
	fmt.Println(pi)
	const name = "Golang"
	fmt.Println(name)
	//name = "Java"
	// this will throw error as constant value cannot be changed

	const (
		port = 9000
		host = "Constant Block"
		c    = 5.67
	)
	fmt.Println(port)
	fmt.Println(host)
	fmt.Println(c)

}
