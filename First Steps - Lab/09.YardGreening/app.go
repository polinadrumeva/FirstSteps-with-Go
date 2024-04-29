package main

import "fmt"

func main() {
	priceForOneM2 := 7.61
	discount := 0.18

	var area float64
	fmt.Scanln(&area)

	priceWithoutDiscount := priceForOneM2 * area
	discountOfPrice := priceWithoutDiscount * discount
	finalPrice := priceWithoutDiscount - discountOfPrice

	fmt.Printf("The final price is: %f lv.\n", finalPrice)
	fmt.Printf("The discount is: %f lv.", discountOfPrice)
}
