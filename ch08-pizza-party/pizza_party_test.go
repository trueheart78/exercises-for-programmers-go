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
