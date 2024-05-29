package main

import "fmt"

func main() {
	var age float64
	var sex string
	fmt.Scanln(&age)
	fmt.Scanln(&sex)

	if sex == "m" && age >= 16 {
		fmt.Println("Mr.")
	} else if sex == "m" && age < 16 {
		fmt.Println("Master")
	} else if sex == "f" && age >= 16 {
		fmt.Println("Ms.")
	} else if sex == "f" && age < 16 {
		fmt.Println("Miss")
	}
}
