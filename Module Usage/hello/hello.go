package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Jack")
	fmt.Println(message)
}
