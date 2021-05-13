package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// msg := greetings.SayHi("nidhal")
	// fmt.Println(msg)

	names := []string{"nidhal", "katharina", "philana"}
	messages, err := greetings.Hellos(names)
	if err == nil {
		fmt.Println(messages)
	}
}
