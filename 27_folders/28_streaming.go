package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sourceFile, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("copy_hello.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}

	}

	writer.Flush()

	fmt.Println("written to new file succesfully")

}
