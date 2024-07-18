package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	sum := 0

	for i := 1; i <= n; i++ {
		var number int
		fmt.Scanln(&number)
		sum += number

	}

	fmt.Println(sum)
}
