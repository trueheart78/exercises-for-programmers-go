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

func (p pizzaParty) totalSlices() int {
	return p.pizzas * p.pizzaSize
}

func (p pizzaParty) guestSlices() int {
	guestSlices := float64(p.totalSlices()) / float64(p.guests)
	return int(math.Floor(guestSlices))
}

func (p pizzaParty) remainingSlices() int {
	return p.totalSlices() - (p.guestSlices() * p.guests)
}

func (p pizzaParty) leftovers() bool {
	if p.remainingSlices() > 0 {
		return true
	}
	return false
}

func pizzaSlices(number int) string {
	slices := []string{}
	for i := 0; i < number; i++ {
		slices = append(slices, " 🍕")
	}
	return strings.Join(slices, "")
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
	fmt.Println(pizzaSlices(party.guestSlices()))
	if party.leftovers() == true {
		fmt.Print("Leftovers will be")
		fmt.Println(pizzaSlices(party.remainingSlices()))
	} else {
		fmt.Println("There will be no leftovers.")
	}
}
