package main

import (
	"fmt"
	"math"
)

func main() {

	var name string
	var timeSerial, timeBreak int
	fmt.Scanln(&name)
	fmt.Scanln(&timeSerial)
	fmt.Scanln(&timeBreak)

	timeLunch := float64(timeBreak / 8)
	timeRest := float64(timeBreak / 4)

	total := timeLunch + timeRest + float64(timeSerial)
	var diff int

	if total <= float64(timeBreak) {
		diff = int(math.Ceil(float64(timeBreak) - float64(total)))
		fmt.Printf("You have enough time to watch %s and left with %d minutes free time.", name, diff)
	} else {
		diff = int(math.Ceil(float64(total)-float64(timeBreak)) + 1)
		fmt.Printf("You don't have enough time to watch %s, you need %d more minutes.", name, diff)
	}
}
