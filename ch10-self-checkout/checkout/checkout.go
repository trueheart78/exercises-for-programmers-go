package checkout

import (
	"strconv"
)

const TaxRate = 5.5

type Register struct {
	items []Item
}

type Item struct {
	name  string
	cents int
}

func newItem(name, cost string) Item {
	parsedCost, err := strconv.ParseFloat(cost, 64)
	var cents int
	if err == nil {
		cents = int(parsedCost * 100)
	}
	return Item{name, cents}
}

func (r Register) itemCount() int {
	return len(r.items)
}

func (r *Register) AddItem(name, cost string) bool {
	item := newItem(name, cost)
	if item.cents != 0 {
		r.items = append(r.items, item)
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
	for _, item := range r.items {
		sum = sum + item.cents
	}
	sum = int((1.0 + (TaxRate / 100)) * float64(sum))
	return sum
}
