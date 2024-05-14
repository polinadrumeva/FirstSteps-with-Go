package main

import "fmt"

func main() {
	var num int
	var bonusPoint float64
	fmt.Scanln(&num)

	if num <= 100 {
		bonusPoint += 5
	} else if num > 100 && num < 1000 {
		bonusPoint += float64(num) * 0.20
	} else if num > 1000 {
		bonusPoint += float64(num) * 0.10
	}

	if num%2 == 0 {
		bonusPoint += 1
	} else if num%5 == 0 {
		bonusPoint += 2
	}

	sum := float64(num) + bonusPoint

	fmt.Println(bonusPoint)
	fmt.Println(sum)

}
