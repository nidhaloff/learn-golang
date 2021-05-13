package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {

	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	msg, err := greetings.Hello("nidhal")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msg)
	}
}
