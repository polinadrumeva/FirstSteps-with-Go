package main

import "fmt"

func main() {
	var day string
	fmt.Scanln(&day)

	if day == "Monday" || day == "Tuesday" || day == "Friday" {
		fmt.Println("12")
	} else if day == "Wednesday" || day == "Thursday" {
		fmt.Println("14")
	} else {
		fmt.Println("16")
	}
}
