package main

import "testing"

func TestSqftPerGallon(t *testing.T) {
	if sqftPerGallon != 350 {
		t.Fatalf("Fail: expected %d, got %d", 350, sqftPerGallon)
	}

	t.Log("OK: const sqftPerGallon")
}

func TestRoomSqft(t *testing.T) {
	room := newRoom("10", "100")
	if room.sqft() != 1000 {
		t.Fatalf("Fail: expected %d, got %d", 1000, room.sqft())
	}

	t.Log("OK: Room.sqft()")
}

func TestRoomGallonsToPaint(t *testing.T) {
	room := newRoom("10", "10")
	if room.gallonsToPaint() != "1" {
		t.Fatalf("Fail: expected %v, got %v", "1", room.gallonsToPaint())
	}

	room = newRoom("10", "100")
	if room.gallonsToPaint() != "3" {
		t.Fatalf("Fail: expected %v, got %v", "3", room.gallonsToPaint())
	}

	t.Log("OK: Room.gallonsToPaint()")
}
