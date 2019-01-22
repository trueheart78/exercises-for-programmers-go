package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var principal = 1500.00
var rate = 4.3
var years = 4.0

func TestNewInterest(t *testing.T) {
	assert := assert.New(t)
	i := NewInterest(principal, rate, years)
	assert.Equal(i, Interest{principal: principal, rate: rate, years: years})
}

func TestInterestAmount(t *testing.T) {
	assert := assert.New(t)
	i := NewInterest(principal, rate, years)
	assert.Equal(1758.00, i.Amount(), "amount total")
}

func TestPercentageRate(t *testing.T) {
	assert := assert.New(t)
	i := NewInterest(principal, rate, years)
	assert.Equal(0.043, i.percentageRate(), "percentage rate")
}
