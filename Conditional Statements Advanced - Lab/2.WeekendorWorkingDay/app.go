package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	if input == "Monday" || input == "Tuesday" || input == "Wednesday" || input == "Thursday" || input == "Friday" {
		fmt.Println("Working day")
	} else if input == "Saturday" || input == "Sunday" {
		fmt.Println("Weekend")
	} else {
		fmt.Println("Error")
	}
}
