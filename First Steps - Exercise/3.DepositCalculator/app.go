package main

import "fmt"

func main() {
	var deposit, months, percent float64
	fmt.Scanln(&deposit)
	fmt.Scanln(&months)
	fmt.Scanln(&percent)
	result := deposit + float64(months)*((deposit*(percent/100))/12)

	fmt.Println(result)
}
