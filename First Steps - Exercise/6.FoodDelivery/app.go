package main

import "fmt"

func main() {
	chickenMenu := 10.35
	fishMenu := 12.40
	vegieMenu := 8.15
	take := 2.50

	var numChickenMenu, numFishMenu, numVegieMenu float64
	fmt.Scanln(&numChickenMenu)
	fmt.Scanln(&numFishMenu)
	fmt.Scanln(&numVegieMenu)
	sumWithoutTake := (chickenMenu * numChickenMenu) + (fishMenu * numFishMenu) + (vegieMenu * numVegieMenu)
	deserts := sumWithoutTake * 0.20

	sum := sumWithoutTake + deserts + take

	fmt.Println(sum)
}
