package main

import "fmt"

func main() {

	var duljina, shirina, visochina, percent float64
	fmt.Scanln(&duljina)
	fmt.Scanln(&shirina)
	fmt.Scanln(&visochina)
	fmt.Scanln(&percent)

	volume := duljina * shirina * visochina
	volumeInL := volume / 1000
	litres := volumeInL * (1 - (percent / 100))

	fmt.Println(litres)
}
