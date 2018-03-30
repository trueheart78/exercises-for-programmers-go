package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const sqftPerGallon = 350

type Room struct {
	width, length int
}

func (r Room) sqft() int {
	return r.width * r.length
}

func (r Room) gallonsToPaint() string {
	sqft := float64(r.sqft())
	gallons := math.Ceil(sqft / sqftPerGallon)
	return strconv.Itoa(int(gallons))
}

func newRoom(width, length string) Room {
	r := Room{}
	r.width, _ = strconv.Atoi(strings.TrimSpace(width))
	r.length, _ = strconv.Atoi(strings.TrimSpace(length))
	return r
}

func main() {
	fmt.Println("🖍  Paint Calc 🖍")
	var width, length string
	fmt.Print("What is the room's width, feet-wise? ")
	fmt.Scanln(&width)
	fmt.Print("What is the room's length, feet-wise? ")
	fmt.Scanln(&length)
	room := newRoom(width, length)
	fmt.Println("It will take", room.gallonsToPaint(), "gallons of paint to cover that", room.sqft(), "square foot ceiling")
}
