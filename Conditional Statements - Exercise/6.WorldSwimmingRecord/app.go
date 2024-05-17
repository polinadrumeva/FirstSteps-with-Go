package main

import (
	"fmt"
	"math"
)

func main() {
	var record, distance, timeForM float64
	fmt.Scanln(&record)
	fmt.Scanln(&distance)
	fmt.Scanln(&timeForM)

	tenp := 12.50
	times := math.Floor(distance / 15)

	sum := (distance * timeForM) + (times * tenp)

	if sum < record {
		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.", sum)
	} else {
		fmt.Printf("No, he failed! He was %.2f seconds slower.", sum-record)
	}
}
