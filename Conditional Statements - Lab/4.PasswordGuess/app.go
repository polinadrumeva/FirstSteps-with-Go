package main

import "fmt"

func main() {
	var password string
	pass := "s3cr3t!P@ssw0rd"
	fmt.Scanln(&password)
	if password == pass {
		fmt.Println("Welcome")
	} else {
		fmt.Println("Wrong password!")
	}
}
