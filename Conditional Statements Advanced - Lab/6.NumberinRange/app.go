package main

import "fmt"

func main() {
	var num float64
	fmt.Scanln(&num)

	if num >= -100 && num < 0 || num > 0 && num <= 100 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
