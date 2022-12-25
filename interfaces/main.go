package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type swedishBot struct{}

func main() {
	fmt.Println("Interface usage demo")
	eb := englishBot{}
	sb := swedishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (swedishBot) getGreeting() string {
	return "Hej"
}
