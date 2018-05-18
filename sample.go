package main

import (
	"fmt"
)

func red() {
	fmt.Println("Red")
}

func green() {
	fmt.Println("Green")
}

var funcMap = map[string]interface{}{
	"red":   red,
	"green": green,
}

func main() {
	fmt.Println("Go")
	funcMap["red"].(func())()
	funcMap["green"].(func())()
}
