package main

import (
	"fmt"
	"github.com/trueheart78/exercises-for-programmers-go/ch10-self-checkout/checkout"
	"strconv"
)

func main() {
	register := checkout.Register{}
	for i := 0; i < 3; i++ {
		var name, cost, qty string
		fmt.Print("What is the name of the item? ")
		fmt.Scanln(&name)

		fmt.Print("What is the cost of the item? $")
		fmt.Scanln(&cost)

		var n int
		var err error
		for {
			fmt.Print("What is the qty of the item? ")
			fmt.Scanln(&qty)

			n, err = strconv.Atoi(qty)
			if err == nil {
				break
			}
		}

		for j := 0; j < n; j++ {
			register.AddItem(name, cost)
		}
	}

	fmt.Println("Total cost (with", checkout.TaxRate, "% tax) is $", register.OutputCost())
}
