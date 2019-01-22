package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/trueheart78/exercises-for-programmers-go/ch12-computing-simple-interest/converter"
)

func main() {
	var principal string
	var rate string
	var years string

	fmt.Println("💰  Interest Caluclator 💰")

	fmt.Print("How much would you like to invest? $")
	fmt.Scanln(&principal)

	validPrincipal, err := strconv.ParseFloat(principal, 64)
	if err != nil {
		fmt.Println("You entered an invalid principal. ")
		os.Exit(1)
	}

	fmt.Print("What is the interest rate? ")
	fmt.Scanln(&rate)

	validRate, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		fmt.Println("You entered an invalid rate.")
		os.Exit(1)
	}

	fmt.Print("How many years are you looking to invest? ")
	fmt.Scanln(&years)

	validYears, err := strconv.ParseFloat(years, 64)
	if err != nil {
		fmt.Println("You entered an invalid number of years.")
		os.Exit(1)
	}

	i := converter.NewInterest(validPrincipal, validRate, validYears)
	fmt.Printf("Your ROI will be $%.2f\n", i.Amount())
	fmt.Println("❤️  Have A Nice Day! ❤️")
}
