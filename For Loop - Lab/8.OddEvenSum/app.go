package main

import (
	"fmt"
	"math"
)

func main() {
	var n, oddSum, evenSum, currentNum int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&currentNum)
		if i%2 == 1 {
			oddSum += currentNum
		} else {
			evenSum += currentNum
		}
	}

	if oddSum == evenSum {
		fmt.Println("Yes")
		fmt.Printf("Sum = %d", oddSum)
	} else {
		fmt.Println("No")
		diff := math.Abs(float64(oddSum) - float64(evenSum))
		fmt.Printf("Diff = %.0f", diff)
	}
}
