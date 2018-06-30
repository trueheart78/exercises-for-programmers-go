package main

import (
	"fmt"
	"strconv"

	"github.com/trueheart78/exercises-for-programmers-go/ch11-currency-converter/converter"
)

func main() {
	var amount string
	fmt.Println("💰  Currency Converter 💰")
	fmt.Printf("Current EURO to USD conversion rate: %.5f\n", converter.USDExchangeRate)

	fmt.Print("How many euro do you want to convert? ")
	fmt.Scanln(&amount)

	euroCents, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("You entered an invalid amount.")
	} else {
		fmt.Printf("Your %.2f in euros converts to $%.2f\n", euroCents, converter.EuroToUSD(euroCents))
	}
	fmt.Println("❤️  Have A Nice Day! ❤️")
}
