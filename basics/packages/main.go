package main

import (
	"fmt"
	"math/rand"

	"rsc.io/quote"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(quote.Glass())
}
