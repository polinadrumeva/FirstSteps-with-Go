package main

import (
	"fmt"
	"math"
)

func main() {
	var kind string
	fmt.Scanln(&kind)
	if kind == "square" {
		var a float64
		fmt.Scanln(&a)
		fmt.Println(a * a)
	} else if kind == "rectangle" {
		var a, b float64
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		fmt.Println(a * b)
	} else if kind == "circle" {
		var a float64
		fmt.Scanln(&a)
		fmt.Println((a * a) * math.Pi)
	} else if kind == "triangle" {
		var a, b float64
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		fmt.Println((a * b) / 2)
	}
}
