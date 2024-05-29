package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)

	if input == 1 {
		fmt.Println("Monday")
	} else if input == 2 {
		fmt.Println("Tuesday")
	} else if input == 3 {
		fmt.Println("Wednesday")
	} else if input == 4 {
		fmt.Println("Thursday")
	} else if input == 5 {
		fmt.Println("Friday")
	} else if input == 6 {
		fmt.Println("Saturday")
	} else if input == 7 {
		fmt.Println("Sunday")
	} else {
		fmt.Println("Error")
	}
}
