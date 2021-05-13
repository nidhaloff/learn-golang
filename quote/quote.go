package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	q := quote.Glass()
	fmt.Println(q)
}