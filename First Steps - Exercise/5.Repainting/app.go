package main

import "fmt"

func main() {
	nailonPk := 1.50
	paintPk := 14.50
	foamL := 5.00

	var pens, markers, foam, discount int
	fmt.Scanln(&pens)
	fmt.Scanln(&markers)
	fmt.Scanln(&foam)
	fmt.Scanln(&discount)
	sum := (penPkg * float64(pens)) + (markersPkg * float64(markers)) + (foamL * float64(foam))
	sumWithDiscount := sum - (sum * (float64(discount) / 100))

	fmt.Println(sumWithDiscount)
}
