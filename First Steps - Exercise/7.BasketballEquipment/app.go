package main

import "fmt"

func main() {

	var taxForYear float64
	fmt.Scanln(&taxForYear)

	shoes := taxForYear - (taxForYear * 0.40)
	clothes := shoes - (shoes * 0.20)
	ball := clothes / 4
	akssessoires := ball / 5

	sum := shoes + clothes + ball + akssessoires + taxForYear

	fmt.Println(sum)
}
