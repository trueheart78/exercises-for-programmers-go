package checkout

import (
	"strconv"
)

const TaxRate = 5.5

type Register struct {
	Items []Item
}

type Item struct {
	Name            string
	Cents, Quantity int
}

func (i Item) Cost() float64 {
	return float64(i.Cents) / 100
}

func newItem(name, cost, qty string) Item {
	parsedCost, err := strconv.ParseFloat(cost, 64)
	var cents, quantity int
	if err == nil {
		cents = int(parsedCost * 100)
	}
	parsedQty, err := strconv.Atoi(qty)
	if err == nil {
		quantity = parsedQty
	}
	return Item{name, cents, quantity}
}

func (r Register) itemCount() int {
	return len(r.Items)
}

func (r *Register) AddItem(name, cost, qty string) bool {
	item := newItem(name, cost, qty)
	if item.Cents != 0 && item.Quantity != 0 {
		r.Items = append(r.Items, item)
		return true
	}
	return false
}

func (r Register) OutputCost() string {
	cost := float64(r.totalCost()) / 100
	return strconv.FormatFloat(cost, 'f', -1, 64)
}

func (r Register) totalCost() int {
	var sum int
	for _, item := range r.Items {
		sum = sum + (item.Cents * item.Quantity)
	}
	sum = int((1.0 + (TaxRate / 100)) * float64(sum))
	return sum
}
