package main

import "fmt"

func main() {
	var town string
	var sells float64
	fmt.Scanln(&town)
	fmt.Scanln(&sells)

	if town == "Sofia" {
		if sells >= 0 && sells <= 500 {
			sum := sells * 0.05
			fmt.Printf("%.2f", sum)
		} else if sells > 500 && sells <= 1000 {
			sum := sells * 0.07
			fmt.Printf("%.2f", sum)
		} else if sells > 1000 && sells <= 10000 {
			sum := sells * 0.08
			fmt.Printf("%.2f", sum)
		} else if sells > 10000 {
			sum := sells * 0.12
			fmt.Printf("%.2f", sum)
		} else if sells < 0 {
			fmt.Println("error")
		}

	} else if town == "Plovdiv" {
		if sells >= 0 && sells <= 500 {
			sum := sells * 0.055
			fmt.Printf("%.2f", sum)
		} else if sells > 500 && sells <= 1000 {
			sum := sells * 0.08
			fmt.Printf("%.2f", sum)
		} else if sells > 1000 && sells <= 10000 {
			sum := sells * 0.12
			fmt.Printf("%.2f", sum)
		} else if sells > 10000 {
			sum := sells * 0.145
			fmt.Printf("%.2f", sum)
		} else if sells < 0 {
			fmt.Println("error")
		}
	} else if town == "Varna" {
		if sells >= 0 && sells <= 500 {
			sum := sells * 0.045
			fmt.Printf("%.2f", sum)
		} else if sells > 500 && sells <= 1000 {
			sum := sells * 0.075
			fmt.Printf("%.2f", sum)
		} else if sells > 1000 && sells <= 10000 {
			sum := sells * 0.10
			fmt.Printf("%.2f", sum)
		} else if sells > 10000 {
			sum := sells * 0.13
			fmt.Printf("%.2f", sum)
		} else if sells < 0 {
			fmt.Println("error")
		}
	} else {
		fmt.Println("error")
	}

}
