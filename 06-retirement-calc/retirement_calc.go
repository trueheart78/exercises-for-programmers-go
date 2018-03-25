package main

import (
	"fmt"
	"strconv"
	"time"
)

type person struct {
	age, retirement_age int
}

func (p *person) assign(age, retirement_age string) {
	p.age, _ = strconv.Atoi(age)
	p.retirement_age, _ = strconv.Atoi(retirement_age)
}

func (p person) yearsToRetirement() int {
	return p.retirement_age - p.age
}

func (p person) yearOfRetirement() int {
	return year() + p.yearsToRetirement()
}

func year() int {
	year, _, _ := time.Now().Date()
	return year
}

func main() {
	var age, retirement_age string
	user := person{}

	fmt.Print("What is your current age? ")
	fmt.Scanln(&age)
	fmt.Print("What age do you want to retire? ")
	fmt.Scanln(&retirement_age)

	user.assign(age, retirement_age)
	fmt.Println("You have", user.yearsToRetirement(), "years to go")
	fmt.Println("You can retire in", user.yearOfRetirement())
}
