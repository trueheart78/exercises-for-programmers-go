package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type pizzaParty struct {
	guests, pizzas, pizzaSize int
}

func newPizzaParty(people, pizzas, slicesPer string) pizzaParty {
	p := pizzaParty{}
	p.guests, _ = strconv.Atoi(strings.TrimSpace(people))
	p.pizzas, _ = strconv.Atoi(strings.TrimSpace(pizzas))
	p.pizzaSize, _ = strconv.Atoi(strings.TrimSpace(slicesPer))
	return p
}

func (p pizzaParty) guestSlices() int {
	totalSlices := p.pizzas * p.pizzaSize
	var guestSlices float64
	guestSlices = float64(totalSlices) / float64(p.guests)
	return int(math.Floor(guestSlices))
}

func main() {
	var people, pizzas, slices string
	fmt.Println("🍕 Pizza Party 🍕")
	fmt.Print("How many guests showed up? 👤  ")
	fmt.Scanln(&people)
	fmt.Print("How many pizzas were ordered? ")
	fmt.Scanln(&pizzas)
	fmt.Print("How many slices for each pizza? ")
	fmt.Scanln(&slices)

	party := newPizzaParty(people, pizzas, slices)
	fmt.Print("Each 👤  should get")
	for i := 0; i < party.guestSlices(); i++ {
		fmt.Print(" 🍕")
	}
	fmt.Println()
}
