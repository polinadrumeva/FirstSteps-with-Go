package main

import "fmt"

func main() {
	var name string
	var projectNumber int
	fmt.Scanln(&name)
	fmt.Scanln(&projectNumber)

	oneProjectHours := 3
	result := projectNumber * oneProjectHours

	fmt.Printf("The architect %s will need %d hours to complete %d project/s.", name, result, projectNumber)
}
