package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	if num1 > num2 {
		fmt.Println(num1)
	} else {
		fmt.Println(num2)
	}
}
