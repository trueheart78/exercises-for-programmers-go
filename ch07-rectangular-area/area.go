package main

import (
	"fmt"
	"strconv"
)

const ConversionRate = 0.09290304

type room struct {
	length, width int
}

func (r *room) assign(length, width string) {
	r.length, _ = strconv.Atoi(length)
	r.width, _ = strconv.Atoi(width)
}

func (r room) squareFeet() int {
	return r.length * r.width
}

func (r room) squareMeters() float64 {
	return float64(r.squareFeet()) * ConversionRate
}

func main() {
	var length, width string
	var newRoom room
	fmt.Print("What is the room's length, feet-wise? ")
	fmt.Scanln(&length)
	fmt.Print("What is the room's width, feet-wise? ")
	fmt.Scanln(&width)
	newRoom.assign(length, width)
	fmt.Printf("The square feet of your room is %d\n", newRoom.squareFeet())
	fmt.Printf("The square meters of your room is %.8f\n", newRoom.squareMeters())
}
