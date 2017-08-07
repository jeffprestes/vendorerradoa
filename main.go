package main

import (
	"fmt"

	"github.com/jeffprestes/vendorerradoa/calc"
)

func main() {
	a := 2
	b := 2
	c := 2
	fmt.Println("Qual o delta? ", calc.Delta(a, b, c))
}
