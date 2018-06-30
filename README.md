# Exercises for Programmers, in Golang

Last time it was Ruby. This time, Go is what I will be using.

## Book

[Exercises for Programmers on PragProg.com][e4p]

## Exercises

1. **Saying Hello:** Create a program that prompts for your name and prints a greeting using your name.
   * [Basic Hello](01-say-hello/hello-basic.go)
2. **Character Counter:** Create a program that prompts for an input string and displays output 
that shows the input string and the number of characters the string contains.
   * [Basic Counter](02-character-count/counter-basic.go)

3. **Printing Quotes:** Create a program that prompts for a quote and an author.
Example: _Obi Wan  said, "These are not the droids you are looking for"_
   * [Basic Quote](03-printing-quotes/quote-basic.go)

4. **Mad Libs:** Create a simple mad-lib program that prompts for a noun, a verb, an adverb, and 
an adjective and injects those into a story that you create. Accept a noun, verb, adjective, and adverb.
   * [Basic Mad Lib](04-mad-lib/basic-mad-lib.go)

5. **Simple Math:** Write a program that prompts for two numbers. Print the sum, difference, 
product, and quotient of those numbers as shown in the example output.
   * [Simple Math](05-simple-math/simple_math.go)
   * [Simple Math Tests](05-simple-math/simple_math_test.go)

6. **Retirement Calculator:** Create a program that determines how many years you have left until 
retirement and the year you can retire. It should prompt for your current age and the age you want 
to retire and display the calculation.
   * [Retirement Calculator](06-retirement-calc/retirement_calc.go)
   * [Retirement Calculator Tests](06-retirement-calc/retirement_calc_test.go)

7. **Rectangular Room:** Create a program that calculates the area of a room. Prompt the user for 
the length and width of the room in feet. Then display the area in both square feet and square 
meters, using the following equation:
   ```
   m**m = f**f * 0.09290304
   ```
   * [Rectangular Area](ch07-rectangular-area/area.go)
   * [Rectangular Area Tests](ch07-rectangular-area/area_test.go)

8. **Pizza Party:** Write a program to evenly divide pizzas. Prompt for the number of people, the 
number of pizzas, and the number of slices per pizza. Ensure that the number of pieces comes out 
even. Display the number of pieces of pizza each person should get. If there are leftovers, show 
the number of leftover pieces.
   * [Pizza Party](ch08-pizza-party/pizza_party.go)
   * [Pizza Party Tests](ch08-pizza-party/pizza_party_test.go)

9. **Paint Calculator:** Calculate gallons of paint needed to paint the ceiling of a room. Prompt 
for the length and width, and assume one gallon covers 350 square feet. Display the number of 
gallons needed to paint the ceiling as a whole number.
   * [Paint Calculator](ch09-paint-calculator/paint_calc.go)
   * [Paint Calculator Tests](ch09-paint-calculator/paint_calc_test.go)

10. **Self-Checkout:** Create a simple self-checkout system. Prompt for the prices and quantities 
of three items. Calculate the subtotal of the items. Then calculate the tax using a tax rate of 5.5%.
Print out the line items with the quantity and total, and then print out the subtotal, tax amount, and total.
    * [Checkout Pkg](ch10-self-checkout/checkout/checkout.go)
    * [Checkout Pkg Tests](ch10-self-checkout/checkout/checkout_test.go)
    * [Self Checkout](ch10-self-checkout/self_checkout.go)

11. **Currency Conversion:** Write a program that converts currency. Specifically, convert euros 
to U.S. dollars. Prompt for the amount of money in euros you have, and prompt for the current 
exchange rate of the euro. Print out the new amount in U.S. dollars. The formula for currency 
conversion is `amount_to = (amount_from * rate_from) / rate_to`, where `rate_to` is generally `1.0`.
More details can be found on [Mathinary's Currency Conversion](http://www.mathinary.com/currency_conversion.jsp)
page.
    * [Converter Pkg](ch11-currency-conversion/converter/converter.go)
    * [Converter Pkg Tests](ch11-currency-conversion/converter/converter_test.go)
    * [Currency Conversion](ch11-currency-conversion/currency.go)

[e4p]: https://pragprog.com/book/bhwb/exercises-for-programmers
