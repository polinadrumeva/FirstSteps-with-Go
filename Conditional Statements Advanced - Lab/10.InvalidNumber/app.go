package main

import "fmt"

func main() {
	var num float64
	fmt.Scanln(&num)

	if num >= 100 && num <= 200 || num == 0 {
		fmt.Println()
	} else {
		fmt.Println("invalid")
	}
}
