package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Quote struct {
	author, saying string
}

func (q Quote) statement() {
	q.author = strings.TrimSpace(q.author)
	q.saying = strings.TrimSpace(q.saying)
	fmt.Print(q.author, " said, \"", q.saying, "\"\n")
}

func main() {
	// fmt.Scanln doesn't return a string with spaces how we want, so we use bufio to make it happen.
	reader := bufio.NewReader(os.Stdin)
	quote := Quote{}

	Quote{"Frankie", "Relax"}.statement()

	fmt.Println("Now tell something someone famous once said:")

	// notice that the \n is in single quotes, as bufio panics otherwise
	quote.saying, _ = reader.ReadString('\n')
	fmt.Print("Who was that famous person? ")
	quote.author, _ = reader.ReadString('\n')

	quote.statement()
}
