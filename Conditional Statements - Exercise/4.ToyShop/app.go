package main

import "fmt"

func main() {

	puzzle := 2.60
	doll := 3
	bear := 4.10
	minion := 8.20
	truck := 2
	discount := 0.25
	rent := 0.10

	var price, countPuzzles, countDolls, countBears, countMinions, countTrucks float64
	fmt.Scanln(&price)
	fmt.Scanln(&countPuzzles)
	fmt.Scanln(&countDolls)
	fmt.Scanln(&countBears)
	fmt.Scanln(&countMinions)
	fmt.Scanln(&countTrucks)

	var totalWithoutDiscount, total float64
	totalWithoutDiscount = (countPuzzles * puzzle) + (countDolls * float64(doll)) + (countBears * bear) + (countMinions * minion) + (countTrucks * float64(truck))
	counts := countPuzzles + countDolls + countBears + countMinions + countTrucks

	if counts >= 50 {
		total = totalWithoutDiscount - (totalWithoutDiscount * discount)
		total -= total * rent
	} else {
		total = totalWithoutDiscount - (totalWithoutDiscount * rent)
	}

	if total > price {
		sum := total - price
		fmt.Printf("Yes! %.2f lv left.", sum)
	} else {
		sum := price - total
		fmt.Printf("Not enough money! %.2f lv needed.", sum)
	}
}
