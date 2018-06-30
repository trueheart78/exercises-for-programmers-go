package converter

import (
	"math"
)

const (
	// EuroExchangeRate on 6/30/18
	EuroExchangeRate = 1.16893
	// USDExchangeRate on 6/30/18
	USDExchangeRate = 0.85548
)

// USDToEuro converts the passed amount to Euros
func USDToEuro(amount float64) float64 {
	return convert(amount, USDExchangeRate)
}

// EuroToUSD converts the passed amount to USD
func EuroToUSD(amount float64) float64 {
	return convert(amount, EuroExchangeRate)
}

func convert(amount float64, rate float64) float64 {
	return round((amount * rate) / 1.0)
}

func round(amount float64) float64 {
	return math.Round(amount*100) / 100
}
