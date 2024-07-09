package main

import "fmt"

func main() {
	var input string
	var r, c float64
	fmt.Scanln(&input)
	fmt.Scanln(&r)
	fmt.Scanln(&c)

	if input == "Premiere" {
		price := 12.00
		sum := (r * c) * price
		fmt.Printf("%.2f leva", sum)
	} else if input == "Normal" {
		price := 7.50
		sum := (r * c) * price
		fmt.Printf("%.2f leva", sum)
	} else if input == "Discount" {
		price := 5.00
		sum := (r * c) * price
		fmt.Printf("%.2f leva", sum)
	}
}
