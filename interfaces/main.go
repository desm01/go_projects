package main

import "fmt"

type bot interface {
	getGreeting() string
}
type spanishBot struct{}
type englishBot struct{}

func main() {
	sb := spanishBot{}
	eb := englishBot{}

	printGreeting(sb)
	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func (eb englishBot) getGreeting() string {
	return "Hello there!"
}
