package converter

// Interest class
type Interest struct {
	principal float64
	rate      float64
	years     float64
}

// NewInterest returns a new Interest object with the expected values
func NewInterest(principal float64, rate float64, years float64) (i Interest) {
	i.principal = principal
	i.rate = rate
	i.years = years
	return
}

// Amount uses the "A = P(1 + rt)" formula to calculate actual interest amount
// Example: 1500(1 + (0.043 × 4)) = 1758
func (i Interest) Amount() float64 {
	return i.principal * (1 + (i.percentageRate() * i.years))
}

func (i Interest) percentageRate() float64 {
	return i.rate / 100
}
