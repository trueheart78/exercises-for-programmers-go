package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUSDExchangeRate(t *testing.T) {
	assert.Equal(t, 0.81268, USDExchangeRate)
}

func TestEuroExchangeRate(t *testing.T) {
	assert.Equal(t, 1.2307, EuroExchangeRate)
}
