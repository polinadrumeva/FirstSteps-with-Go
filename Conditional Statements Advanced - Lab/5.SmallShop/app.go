package main

import "fmt"

func main() {
	var product, town string
	var count float64
	fmt.Scanln(&product)
	fmt.Scanln(&town)
	fmt.Scanln(&count)

	if town == "Sofia" {
		if product == "coffee" {
			sum := 0.50 * count
			fmt.Println(sum)
		} else if product == "water" {
			sum := 0.80 * count
			fmt.Println(sum)
		} else if product == "beer" {
			sum := 1.20 * count
			fmt.Println(sum)
		} else if product == "sweets" {
			sum := 1.45 * count
			fmt.Println(sum)
		} else if product == "peanuts" {
			sum := 1.60 * count
			fmt.Println(sum)
		}

	} else if town == "Plovdiv" {
		if product == "coffee" {
			sum := 0.40 * count
			fmt.Println(sum)
		} else if product == "water" {
			sum := 0.70 * count
			fmt.Println(sum)
		} else if product == "beer" {
			sum := 1.15 * count
			fmt.Println(sum)
		} else if product == "sweets" {
			sum := 1.30 * count
			fmt.Println(sum)
		} else if product == "peanuts" {
			sum := 1.50 * count
			fmt.Println(sum)
		}
	} else if town == "Varna" {
		if product == "coffee" {
			sum := 0.45 * count
			fmt.Println(sum)
		} else if product == "water" {
			sum := 0.70 * count
			fmt.Println(sum)
		} else if product == "beer" {
			sum := 1.10 * count
			fmt.Println(sum)
		} else if product == "sweets" {
			sum := 1.35 * count
			fmt.Println(sum)
		} else if product == "peanuts" {
			sum := 1.55 * count
			fmt.Println(sum)
		}
	}

}
