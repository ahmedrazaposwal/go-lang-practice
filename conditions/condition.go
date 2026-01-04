package main

import (
	"fmt"
)

func main() {
	if 15 > 10 {
		fmt.Println("15 is greater than 10")
		fmt.Println(15 > 10)
	}

	time := 14
	if time < 15 {
		fmt.Println("Good day")
	} else {
		fmt.Println("Good evening")
	}

	day := 2

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
		fmt.Println("Unknown")
	}
	fmt.Println()
}
