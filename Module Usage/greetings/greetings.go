package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome from Htet Aung's Greetings module!", name)
	return message
}
