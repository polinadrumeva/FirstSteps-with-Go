package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	if input == "dog" {
		fmt.Println("mammal")
	} else if input == "crocodile" || input == "tortoise" || input == "snake" {
		fmt.Println("reptile")
	} else {
		fmt.Println("unknown")
	}
}
