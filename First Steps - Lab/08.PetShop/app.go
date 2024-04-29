package main

import "fmt"

func main() {
	oneDogFoodPrice := 2.50
	oneCatFoodPrice := 4

	var numDogFood, numCatFood int
	fmt.Scanln(&numDogFood)
	fmt.Scanln(&numCatFood)

	dogFood := oneDogFoodPrice * float64(numDogFood)
	catFood := oneCatFoodPrice * numCatFood
	result := dogFood + float64(catFood)

	fmt.Printf("%f lv.", result)
}
