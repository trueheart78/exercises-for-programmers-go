package main

import "testing"

var subject = room{length, width}
var length = 10
var width = 13
var conversion = 0.09290304

func TestConversionRate(t *testing.T) {
	if ConversionRate != conversion {
		t.Fatalf("Fail: Unexpected ConversionRate value, expected %.8f, got %.8f", conversion, ConversionRate)
	}
	t.Log("OK: ConversionRate")
}

func TestRoomAssign(t *testing.T) {
	new_room := room{}
	new_room.assign("1", "2")
	if new_room.length != 1 {
		t.Fatalf("Fail: room.assign() length, expected %d, got %d", 1, new_room.length)
	}
	if new_room.width != 2 {
		t.Fatalf("Fail: room.assign() width, expected %d, got %d", 2, new_room.width)
	}
	t.Log("OK: room.assign()")
}

func TestRoomSquareFeet(t *testing.T) {
	if subject.squareFeet() != length*width {
		t.Fatalf("Fail: room.squareFeet(), expected %d, got %d", length*width, subject.squareFeet())
	}
	t.Log("OK: room.squareFeet()")
}

func TestRoomSquareMeters(t *testing.T) {
	square_meters := float64(length*width) * conversion
	if subject.squareMeters() != square_meters {
		t.Fatalf("Fail: room.squareMeters(), expected %.8f, got %.8f", square_meters, subject.squareMeters())
	}
	t.Log("OK: room.squareMeters()")
}
