package main

import "fmt"

func main() {
	var firstName, lastName, town string
	fmt.Scanln(&firstName)
	fmt.Scanln(&lastName)
	var age int
	fmt.Scanln(&age)
	fmt.Scanln(&town)

	fmt.Printf("You are %s %s, a %d-years old person from %s.", firstName, lastName, age, town)
}
