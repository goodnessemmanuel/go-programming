package main

import (
	"fmt"

	"github.com/goodnessemmanuel/go-programming/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Oche")
    fmt.Println(message)
}