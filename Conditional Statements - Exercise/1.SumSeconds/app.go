package main

import "fmt"

func main() {
	var first, second, third int
	fmt.Scanln(&first)
	fmt.Scanln(&second)
	fmt.Scanln(&third)

	var min, seconds int
	if first+second >= 60 {
		min += 1
		seconds += (first + second) - 60
	} else {
		seconds = first + second
	}
	if seconds+third >= 60 {
		min += 1
		seconds = (seconds + third) - 60
	} else {
		seconds += third
	}

	if seconds < 10 {
		fmt.Printf("%d:0%d", min, seconds)
	} else {
		fmt.Printf("%d:%d", min, seconds)
	}
}
