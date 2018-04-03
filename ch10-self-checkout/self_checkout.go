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

		var err error
		for {
			fmt.Print("What is the qty of the item? ")
			fmt.Scanln(&qty)

			_, err = strconv.Atoi(qty)
			if err == nil {
				break
			}
		}

		register.AddItem(name, cost, qty)
	}

	fmt.Println("---------")
	for i, item := range register.Items {
		n := i + 1
		fmt.Printf("%d. %v x %d @ %.2f each\n", n, item.Name, item.Quantity, item.Cost())
	}

	fmt.Println("---------")
	fmt.Println("Total cost (with", checkout.TaxRate, "% tax) is $", register.OutputCost())
}
