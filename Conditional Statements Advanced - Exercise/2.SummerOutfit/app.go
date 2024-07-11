package main

import "fmt"

func main() {
	var gradus int
	var input string
	fmt.Scanln(&gradus)
	fmt.Scanln(&input)

	if gradus >= 10 && gradus <= 18 {
		if input == "Morning" {
			fmt.Printf("It's %d degrees, get your Sweatshirt and Sneakers.", gradus)
		} else if input == "Afternoon" {
			fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", gradus)
		} else if input == "Evening" {
			fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", gradus)
		}
	} else if gradus > 18 && gradus <= 24 {
		if input == "Morning" {
			fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", gradus)
		} else if input == "Afternoon" {
			fmt.Printf("It's %d degrees, get your T-Shirt and Sandals.", gradus)
		} else if input == "Evening" {
			fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", gradus)
		}
	} else if gradus >= 25 {
		if input == "Morning" {
			fmt.Printf("It's %d degrees, get your T-Shirt and Sandals.", gradus)
		} else if input == "Afternoon" {
			fmt.Printf("It's %d degrees, get your Swim Suit and Barefoot.", gradus)
		} else if input == "Evening" {
			fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", gradus)
		}
	}
}
