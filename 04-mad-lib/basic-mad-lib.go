package main

import (
	"fmt"
)

type MadLib struct {
	noun, verb, adjective, adverb string
}

func (m MadLib) println() {
	fmt.Println("I can't believe you", m.verb, "your", m.adjective, m.noun, m.adverb)
}

func main() {
	newMadLib := MadLib{}

	fmt.Print("Enter a noun: ")
	fmt.Scanln(&newMadLib.noun)

	fmt.Print("Enter a verb: ")
	fmt.Scanln(&newMadLib.verb)

	fmt.Print("Enter an adjective: ")
	fmt.Scanln(&newMadLib.adjective)

	fmt.Print("Enter an adverb: ")
	fmt.Scanln(&newMadLib.adverb)

	newMadLib.println()
}
