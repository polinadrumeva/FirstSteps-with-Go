package main

import "fmt"

func main() {
	var hour, min int
	fmt.Scanln(&hour)
	fmt.Scanln(&min)
	time := 15
	if min+time >= 59 {
		hour += 1
		min = (min + time) - 59
		if hour > 23 {
			hour = 0
		}
	} else {
		min += time
	}

	if min < 10 {
		fmt.Printf("%d:0%d", hour, min)
	} else {
		fmt.Printf("%d:%d", hour, min)
	}
}
