package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Gimme a string to count! Please! ")
	fmt.Scanln(&name)
	fmt.Println("The number of characters in", name, "is", len(name))
}
