package main

import "fmt"

func main() {
	var budjet, statists, pricePerOne float64
	fmt.Scanln(&budjet)
	fmt.Scanln(&statists)
	fmt.Scanln(&pricePerOne)

	decor := budjet * 0.10
	discount := 0.10

	var sum float64
	if statists > 150 {
		sum = decor + ((statists * pricePerOne) - ((statists * pricePerOne) * discount))
	} else {
		sum = decor + (statists * pricePerOne)
	}

	if sum > budjet {
		fmt.Println("Not enough money!")
		fmt.Printf("Wingard needs %.2f leva more.", sum-budjet)
	} else {
		fmt.Println("Action!")
		fmt.Printf("Wingard starts filming with %.2f leva left.", budjet-sum)
	}
}
