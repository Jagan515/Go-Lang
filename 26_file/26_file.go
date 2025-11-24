package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// get file info
	fileinfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File Name:", fileinfo.Name())
	fmt.Println("File Size:", fileinfo.Size())
	fmt.Println("File Mode:", fileinfo.Mode())
	fmt.Println("File ModTime:", fileinfo.ModTime())
	fmt.Println("Is Directory?:", fileinfo.IsDir())

	// read file content
	data := make([]byte, fileinfo.Size())
	count, err := f.Read(data)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("Bytes read:", count)
	fmt.Println("File Content:\n", string(data))

	// read file using os.ReadFile
	data2, err := os.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File Content using os.ReadFile:\n", string(data2))

	// write to a file
	err = os.WriteFile("output.txt", []byte("Hello, Go File Handling!"), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data written to output.txt successfully")

	// append to a file
	f2, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	_, err = f2.WriteString("\nAppended line.")
	if err != nil {
		panic(err)
	}
	fmt.Println("Data appended to output.txt successfully")

	// delete a file
	err = os.Remove("output.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("output.txt deleted successfully")

	// read folder contents
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	fmt.Println("Directory contents:")
	for _, entry := range entries {
		fmt.Println(" -", entry.Name())
	}

	// create a new folder
	err = os.Mkdir("new_folder", 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("new_folder created successfully")

	// remove the folder
	err = os.Remove("new_folder")
	if err != nil {
		panic(err)
	}
	fmt.Println("new_folder removed successfully")

	// get absolute path
	absPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current Absolute Path:", absPath)

}
