package main

import "fmt"

func main() {
	var product string
	fmt.Scanln(&product)

	if product == "banana" || product == "apple" || product == "kiwi" || product == "cherry" || product == "lemon" || product == "grapes" {
		fmt.Println("fruit")
	} else if product == "tomato" || product == "cucumber" || product == "pepper" || product == "carrot" {
		fmt.Println("vegetable")
	} else {
		fmt.Println("unknown")
	}
}
