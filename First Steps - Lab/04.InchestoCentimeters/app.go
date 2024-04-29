package main

import "fmt"

func main() {

	var inch float64
	fmt.Scanln(&inch)
	floatNum := 2.54

	var cm float64 = inch * floatNum
	fmt.Println(cm)

}
