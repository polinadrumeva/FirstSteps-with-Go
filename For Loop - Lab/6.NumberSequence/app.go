package main

import (
	"fmt"
	"math"
)

func main() {
	var n, currentNum int
	fmt.Scanln(&n)
	minNum := math.MaxInt32
	maxNum := math.MinInt32

	for i := 0; i < n; i++ {
		fmt.Scanln(&currentNum)
		if currentNum < minNum {
			minNum = currentNum
		}

		if currentNum > maxNum {
			maxNum = currentNum
		}

	}

	fmt.Printf("Max number: %d\n", maxNum)
	fmt.Printf("Min number: %d", minNum)
}
