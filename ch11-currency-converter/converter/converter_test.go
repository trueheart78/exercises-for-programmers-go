package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUSDExchangeRate(t *testing.T) {
	assert.Equal(t, 0.85548, USDExchangeRate)
}

func TestEuroExchangeRate(t *testing.T) {
	assert.Equal(t, 1.16893, EuroExchangeRate)
}

func TestConversionFunctions(t *testing.T) {
	conversions := [...][2]float64{
		[2]float64{1.01, 0.86},
		[2]float64{234.56, 200.66},
		[2]float64{1237.89, 1058.99},
	}

	var usd, euro float64
	for _, c := range conversions {
		usd = c[0]
		euro = c[1]

		assert.Equal(t, euro, USDToEuro(usd))
		assert.Equal(t, usd, EuroToUSD(euro))
	}
}
