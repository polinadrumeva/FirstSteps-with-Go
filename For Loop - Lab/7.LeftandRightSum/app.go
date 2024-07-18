package main

import (
	"fmt"
	"math"
)

func main() {
	var n, leftSum, rightSum, currentNum int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&currentNum)
		leftSum += currentNum
	}

	for i := 1; i <= n; i++ {
		fmt.Scanln(&currentNum)
		rightSum += currentNum
	}

	if leftSum == rightSum {
		fmt.Printf("Yes, sum = %d", leftSum)
	} else {
		diff := math.Abs(float64(leftSum) - float64(rightSum))
		fmt.Printf("No, diff = %.0f", diff)
	}
}
