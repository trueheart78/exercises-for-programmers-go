package main

import (
	"fmt"
	"strconv"
)

type twoInt struct {
	a, b int
}

func (t *twoInt) assign(a, b string) {
	t.a, _ = strconv.Atoi(a)
	t.b, _ = strconv.Atoi(b)
}

func (t twoInt) multiply() int {
	return t.a * t.b
}

func (t twoInt) divide() int {
	return t.a / t.b
}

func (t twoInt) add() int {
	return t.a + t.b
}

func (t twoInt) subtract() int {
	return t.a - t.b
}

func main() {
	ints := twoInt{}
	var a, b string
	fmt.Print("I need a number: ")
	fmt.Scanln(&a)
	fmt.Print("I another number: ")
	fmt.Scanln(&b)

	ints.assign(a, b)

	fmt.Println(ints.a, "*", ints.b, "=", ints.multiply())
	fmt.Println(ints.a, "/", ints.b, "=", ints.divide())
	fmt.Println(ints.a, "+", ints.b, "=", ints.add())
	fmt.Println(ints.a, "-", ints.b, "=", ints.subtract())
}
