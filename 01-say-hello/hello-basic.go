package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scanln(&name)
	if name == "" {
		name = "Mark"
	}
	fmt.Println("Oh, Hi", name)
}
