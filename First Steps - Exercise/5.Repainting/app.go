package main

import "fmt"

func main() {
	nailonPk := 1.50
	paintPk := 14.50
	foamL := 5.00
	bags := 0.40

	var nailon, paint, foam, hours float64
	fmt.Scanln(&nailon)
	fmt.Scanln(&paint)
	fmt.Scanln(&foam)
	fmt.Scanln(&hours)
	sumMaterials := ((nailon + 2) * nailonPk) + ((paint + (paint * 0.10)) * (paintPk)) + (foam * foamL) + bags
	workForHour := sumMaterials * 0.30
	sum := sumMaterials + (workForHour * hours)

	fmt.Println(sum)
}
