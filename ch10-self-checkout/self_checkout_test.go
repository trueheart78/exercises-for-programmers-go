package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	received := "received"
	expected := "expected"
	if received != expected {
		t.Fatalf("Fail: expected %v, got %v", expected, received)
	}
	t.Log("Pass: sample test")
}
