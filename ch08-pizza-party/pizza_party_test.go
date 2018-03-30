package main

import "testing"

func TestNewPizzaParty(t *testing.T) {
	party := newPizzaParty("10", "08", " 6\n")
	if party.guests != 10 {
		t.Fatalf("Fail: newPizzaParty().guests, expected %d, got %d", 10, party.guests)
	}
	if party.pizzas != 8 {
		t.Fatalf("Fail: newPizzaParty().pizzas, expected %d, got %d", 8, party.pizzas)
	}
	if party.pizzaSize != 6 {
		t.Fatalf("Fail: newPizzaParty().pizzaSize, expected %d, got %d", 6, party.pizzaSize)
	}
	t.Log("OK: newPizzaParty()")
}

func TestPizzaPartyTotalSlices(t *testing.T) {
	party := newPizzaParty("10", "10", "10")

	if party.totalSlices() != 100 {
		t.Fatalf("Fail: pizzaParty.totalSlices(), expected %d, got %d", 100, party.totalSlices())
	}
	t.Log("OK: pizzaParty.totalSlices()")
}

func TestPizzaPartyGuestSlices(t *testing.T) {
	// when everything is divisible evenly: (10*10) / 10
	party := newPizzaParty("10", "10", "10")
	if party.guestSlices() != 10 {
		t.Fatalf("Fail: pizzaParty.guestSlices() [even], expected %d, got %d", 10, party.guestSlices())
	}

	// when everything is not divisible evenly: (5*6) / 7
	party = newPizzaParty("7", "5", "6")
	if party.guestSlices() != 4.0 {
		t.Fatalf("Fail: pizzaParty.guestSlices() [uneven], expected %d, got %d", 4, party.guestSlices())
	}
	t.Log("OK: pizzaParty.guestSlices()")
}

func TestPizzaPartyRemainingSlices(t *testing.T) {
	// when everything is divisible evenly: (10*10) / 10
	party := newPizzaParty("10", "10", "10")
	if party.remainingSlices() != 0 {
		t.Fatalf("Fail: pizzaParty.remainingSlices() [even], expected %d, got %d", 0, party.remainingSlices())
	}

	// when everything is not divisible evenly: (5*6) / 7
	party = newPizzaParty("7", "5", "6")
	if party.remainingSlices() != 2 {
		t.Fatalf("Fail: pizzaParty.remainingSlices() [uneven], expected %d, got %d", 2, party.remainingSlices())
	}
	t.Log("OK: pizzaParty.remainingSlices()")
}

func TestPizzaPartLeftovers(t *testing.T) {
	// when everything is divisible evenly: (10*10) / 10
	party := newPizzaParty("10", "10", "10")
	if party.leftovers() != false {
		t.Fatalf("Fail: pizzaParty.leftovers() [even], expected %v, got %v", false, party.leftovers())
	}

	// when everything is not divisible evenly: (5*6) / 7
	party = newPizzaParty("7", "5", "6")
	if party.leftovers() != true {
		t.Fatalf("Fail: pizzaParty.leftovers() [uneven], expected %v, got %v", true, party.leftovers())
	}
	t.Log("OK: pizzaParty.leftovers()")
}

func TestPizzaSlices(t *testing.T) {
	sample := pizzaSlices(5)
	if sample != " 🍕 🍕 🍕 🍕 🍕" {
		t.Fatalf("Fail: pizzaSlices(), expected %v, got %v", "x", sample)
	}
	t.Log("OK: pizzaSlices()")
}
