package main

import "fmt"

func main() {
	penPkg := 5.80
	markersPkg := 7.20
	foamL := 1.20

	var pens, markers, foam, discount int
	fmt.Scanln(&pens)
	fmt.Scanln(&markers)
	fmt.Scanln(&foam)
	fmt.Scanln(&discount)
	sum := (penPkg * float64(pens)) + (markersPkg * float64(markers)) + (foamL * float64(foam))
	sumWithDiscount := sum - (sum * (float64(discount) / 100))

	fmt.Println(sumWithDiscount)
}
