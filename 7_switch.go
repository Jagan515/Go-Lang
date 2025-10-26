package main

import "fmt"
import "time"

func main() {
	fmt.Println("Hello, World!")
	// This is a simple switch case example
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with multiple cases
	month := time.Now().Month()
	switch month {
	case time.December, time.January, time.February:
		fmt.Println("It's Winter!")
	case time.March, time.April, time.May:
		fmt.Println("It's Spring!")
	case time.June, time.July, time.August:
		fmt.Println("It's Summer!")
	case time.September, time.October, time.November:
		fmt.Println("It's Autumn!")
	}

	// Switch without an expression
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// Type switch
	var x interface{}
	x = 42
	switch v := x.(type) {
	case int:
		// 'v' is of type int here
		fmt.Println("x is an integer:", v)
	case string:
		fmt.Println("x is a string:", v)
	default:
		fmt.Println("Unknown type")
	}

}
