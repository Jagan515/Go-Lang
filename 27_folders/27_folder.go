package main

import (
	"fmt"
	"os"
	"time"
)

// read folder contents
func main() {
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

	// wait for 1 seconds
	time.Sleep(1 * time.Second)

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
