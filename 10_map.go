package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("Hello, Maps!")

	// declare and initialize a map

	// we use map[keyType]valueType{}

	myMap := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 15,
	}
	fmt.Println("Initial map:", myMap)

	// add a new key-value pair
	myMap["grape"] = 20
	fmt.Println("Map after adding grape:", myMap)

	// access a value by key
	bananaCount := myMap["banana"]
	fmt.Println("Number of bananas:", bananaCount)

	// iterate over the map
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// check if a key exists
	if val, exists := myMap["orange"]; exists {
		fmt.Println("Orange exists with value:", val)
	} else {
		fmt.Println("Orange does not exist")
	}

	// delete a key-value pair
	delete(myMap, "apple")
	fmt.Println("Map after deleting apple:", myMap)

	// get the length of the map
	fmt.Println("Length of map:", len(myMap))

	// demonstrate map with struct values
	type Person struct {
		Name string
		Age  int
	}

	personMap := make(map[string]Person)
	personMap["john"] = Person{Name: "John Doe", Age: 30}
	personMap["jane"] = Person{Name: "Jane Smith", Age: 25}

	fmt.Println("Person map:", personMap)

	// access struct fields from the map
	john := personMap["john"]
	fmt.Printf("John's Age: %d\n", john.Age)

	person := personMap["jane"]
	fmt.Printf("Jane's Name: %s\n", person.Name)

	// maps equals

	anotherMap := map[string]int{
		"banana": 10,
		"orange": 15,
		"grape":  20,
	}

	fmt.Println(maps.Equal(myMap, anotherMap)) // true

}
