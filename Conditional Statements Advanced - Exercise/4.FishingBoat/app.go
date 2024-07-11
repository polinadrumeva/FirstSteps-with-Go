package main

import "fmt"

func main() {
	var budjet int
	var season string
	var num int

	fmt.Scanln(&budjet)
	fmt.Scanln(&season)
	fmt.Scanln(&num)

	var sum float64

	if season == "Spring" {
		if num <= 6 {
			sum = 3000 * 0.90
		} else if num > 6 && num <= 11 {
			sum = 3000 * 0.85
		} else {
			sum = 3000 * 0.75
		}
	} else if season == "Summer" || season == "Autumn" {
		if num <= 6 {
			sum = 4200 * 0.90
		} else if num > 6 && num <= 11 {
			sum = 4200 * 0.85
		} else {
			sum = 4200 * 0.75
		}
	} else if season == "Winter" {
		if num <= 6 {
			sum = 2600 * 0.90
		} else if num > 6 && num <= 11 {
			sum = 2600 * 0.85
		} else {
			sum = 2600 * 0.75
		}
	}

	if num%2 == 0 && season != "Autumn" {
		sum = sum * 0.95
	}

	if float64(budjet) > sum {
		fmt.Printf("Yes! You have %.2f leva left.", float64(budjet)-sum)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva.", sum-float64(budjet))
	}

}
