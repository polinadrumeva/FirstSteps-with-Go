package main

import (
	"fmt"
	"math"
)

func main() {
	var radians float64
	fmt.Scanln(&radians)
	result := radians * 180 / math.Pi

	fmt.Println(result)
}
