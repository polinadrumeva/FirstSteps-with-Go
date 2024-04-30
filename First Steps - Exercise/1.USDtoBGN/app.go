package main

import "fmt"

func main() {
	var usd float64
	fmt.Scanln(&usd)
	fixCourse := 1.79549
	result := usd * fixCourse

	fmt.Println(result)
}
