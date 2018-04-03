package checkout

import (
	"testing"
)

var name = "name"
var qty = "1"
var cost = "12.99"
var costWithTax = "13.7"
var cents = 1299
var totalWithTax = 1370

func TestTaxRate(t *testing.T) {
	if TaxRate != 5.5 {
		t.Fatalf("Fail: expected %.1f, got %.1f", 5.5, TaxRate)
	}
	t.Log("Pass: TestTaxRate()")
}

func TestNewItem(t *testing.T) {
	item := newItem(name, cost, qty)
	if item.Name != name {
		t.Fatalf("Fail: [name] expected %v, got %v", name, item.Name)
	}
	if item.Cents != cents {
		t.Fatalf("Fail: [cents] expected %d, got %d", cents, item.Cents)
	}
	if item.Quantity != 1 {
		t.Fatalf("Fail: [quantity] expected %d, got %d", 1, item.Quantity)
	}
	item = newItem(name, "bad", qty)
	if item.Cents != 0 {
		t.Fatalf("Fail: expected %d, got %d", 0, item.Cents)
	}
	item = newItem(name, cost, "bad")
	if item.Quantity != 0 {
		t.Fatalf("Fail: expected %d, got %d", 0, item.Quantity)
	}
	t.Log("Pass: TestNewItem()")
}

func TestRegisterAddItem(t *testing.T) {
	register := Register{}
	// with a valid item
	success := register.AddItem(name, cost, qty)
	if success != true {
		t.Fatalf("Fail: valid item, expected %v, got %v", true, success)
	}
	if len(register.Items) != 1 {
		t.Fatalf("Fail: valid item, expected %d, got %d", 1, len(register.Items))
	}

	// invalid item
	success = register.AddItem(name, "bad", qty)
	if success != false {
		t.Fatalf("Fail: invalid item, expected %v, got %v", false, success)
	}
	if len(register.Items) != 1 {
		t.Fatalf("Fail: invalid item, expected %d, got %d", 1, len(register.Items))
	}

	t.Log("Pass: TestRegisterAddItem()")
}

func TestRegisterItemCount(t *testing.T) {
	register := Register{}
	register.AddItem(name, cost, qty)
	register.AddItem(name, cost, qty)

	if register.itemCount() != 2 {
		t.Fatalf("Fail: expected %d, got %d", 1, register.itemCount())
	}
	t.Log("Pass: TestRegisterItemCount()")
}

func TestRegisterTotalCost(t *testing.T) {
	register := Register{}
	register.AddItem(name, cost, qty)

	if register.totalCost() != totalWithTax {
		t.Fatalf("Fail: expected %d, got %d", totalWithTax, register.totalCost())
	}
	t.Log("Pass: TestRegisterTotalCost()")
}

func TestRegisterOutputCost(t *testing.T) {
	register := Register{}
	register.AddItem(name, cost, qty)

	if register.OutputCost() != costWithTax {
		t.Fatalf("Fail: expected %v, got %v", costWithTax, register.OutputCost())
	}
	t.Log("Pass: TestRegisterOutputCost()")
}
