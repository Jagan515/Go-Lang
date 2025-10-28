package main

import (
	"fmt"
	"time"
)

// Step 1: Define the Person struct
type Person struct {
	name      string
	age       int64
	createdAt time.Time // stores creation time
}

// Step 2: Function to extract time details from a time.Time value
func extractTime(t time.Time) {
	hr := t.Hour()
	min := t.Minute()
	sec := t.Second()
	fmt.Printf("Extracted Time: %02d:%02d:%02d\n", hr, min, sec)
}

// Attaching Behavior to Structs using Methods (Optional)
// You can define methods on structs to give them behavior.
// For example, a method to display person's info could be added here.
// convenience method to display person info p
func (p Person) displayInfo() {
	fmt.Printf("Name: %s, Age: %d, Created At: %s\n", p.name, p.age, p.createdAt.String())
}

// modfiing the struct field using pointer receiver
func (p *Person) haveBirthday() {
	p.age += 1
}

// constructor function (Optional)
func NewPerson(name string, age int64) Person {
	return Person{name: name, age: age, createdAt: time.Now()}
}

// similar to above constructor but returns pointer to struct
func NewPersonPointer(name string, age int64) *Person {
	return &Person{name: name, age: age, createdAt: time.Now()}
}

func main() {
	// Step 3: Create Person p1 without creation time
	p1 := Person{name: "Alice", age: 30, createdAt: time.Now()}
	fmt.Println("Person 1:", p1)

	// Step 4: Create Person p2 with creation time
	p2 := Person{"Bob", 25, time.Now()}
	fmt.Println("Person 2:", p2)

	// Step 5: Access and modify struct fields
	fmt.Println("Name of p1:", p1.name)
	fmt.Println("Age of p2:", p2.age)

	p1.age = 31
	fmt.Println("Updated age of p1:", p1.age)

	// Step 6: Use extractTime function to show p2's creation time
	extractTime(p2.createdAt)
	// Optional: Display person's info using method
	p1.displayInfo()
	p2.displayInfo()
	// Modifying struct field using pointer receiver method
	p1.haveBirthday()
	fmt.Println("After birthday, p1's age:", p1.age)

	// Using constructor functions
	p3 := NewPerson("Charlie", 28)
	fmt.Println("Person 3 (using constructor):", p3)

	p4 := NewPersonPointer("Diana", 22)
	fmt.Println("Person 4 (using pointer constructor):", *p4)

	p3.displayInfo()
	p4.displayInfo()

	order := struct {
		id     int
		amount float64
	}{
		id:     101,
		amount: 250.75,
	}
	fmt.Println("Order:", order)

}
