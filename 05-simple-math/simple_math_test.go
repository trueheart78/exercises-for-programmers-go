package main

import "testing"

var subject = twoInt{10, 20}

func TestMultiply(t *testing.T) {
	if subject.multiply() != subject.a*subject.b {
		t.Fatal("Fail: twoInt.multiply()")
	}
	t.Log("OK: twoInt.multiply()")
}

func TestDivide(t *testing.T) {
	if subject.divide() != subject.a/subject.b {
		t.Fatal("Fail: twoInt.divide()")
	}
	t.Log("OK: twoInt.divide()")
}

func TestAdd(t *testing.T) {
	if subject.add() != subject.a+subject.b {
		t.Fatal("Fail: twoInt.add()")
	}
	t.Log("OK: twoInt.add()")
}

func TestSubtract(t *testing.T) {
	if subject.subtract() != subject.a-subject.b {
		t.Fatal("Fail: twoInt.subtract()")
	}
	t.Log("OK: twoInt.subtract()")
}
