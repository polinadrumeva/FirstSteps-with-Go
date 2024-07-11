package main

import "fmt"

func main() {
	roses := 5
	dalias := 3.8
	tulips := 2.8
	narcis := 3
	gladiols := 2.5

	discountRoses := 0.90
	discountDaliasTulips := 0.85

	var input string
	var num int
	var budjet float64
	fmt.Scanln(&input)
	fmt.Scanln(&num)
	fmt.Scanln(&budjet)

	var sum float64

	if input == "Roses" {
		if num > 80 {
			sum = (float64(num) * float64(roses)) * discountRoses
		} else {
			sum = float64(num) * float64(roses)
		}
	} else if input == "Dahlias" {
		if num > 90 {
			sum = (float64(num) * float64(dalias)) * discountDaliasTulips
		} else {
			sum = float64(num) * float64(dalias)
		}
	} else if input == "Tulips" {
		if num > 80 {
			sum = (float64(num) * float64(tulips)) * discountDaliasTulips
		} else {
			sum = float64(num) * float64(tulips)
		}
	} else if input == "Narcissus" {
		if num < 120 {
			sum = (float64(num) * float64(narcis)) * 1.15
		} else {
			sum = float64(num) * float64(narcis)
		}
	} else if input == "Gladiolus" {
		if num < 80 {
			sum = (float64(num) * float64(gladiols)) * 1.20
		} else {
			sum = float64(num) * float64(gladiols)
		}
	}

	if budjet > sum {
		fmt.Printf("Hey, you have a great garden with %d %s and %.2f leva left.", num, input, budjet-sum)
	} else {
		fmt.Printf("Not enough money, you need %.2f leva more.", sum-budjet)
	}

}
