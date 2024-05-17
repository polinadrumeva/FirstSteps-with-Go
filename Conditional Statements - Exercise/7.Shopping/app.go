package main

import "fmt"

func main() {

	var budjet, numVideo, numProc, numRam float64
	fmt.Scanln(&budjet)
	fmt.Scanln(&numVideo)
	fmt.Scanln(&numProc)
	fmt.Scanln(&numRam)

	videocarts := 250
	processor := (numVideo * float64(videocarts)) * 0.35
	ram := (numVideo * float64(videocarts)) * 0.10

	var sum float64
	if numVideo > numProc {
		sum = ((numVideo * float64(videocarts)) + (numProc * processor) + (numRam * ram)) * 0.85
	} else {
		sum = (numVideo * float64(videocarts)) + (numProc * processor) + (numRam * ram)
	}

	if sum < budjet {
		fmt.Printf("You have %.2f leva left!", budjet-sum)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva more!", sum-budjet)
	}
}
