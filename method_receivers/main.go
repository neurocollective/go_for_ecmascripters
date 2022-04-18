package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
}

func (p *Person) Introduce() {
	personObject := *p
	first := personObject.FirstName
	last := personObject.LastName
	fmt.Println("Hello, I am", first, last)
}

func main() {

	dude := Person { "Cody", "McCoder"}

	dude.Introduce()
}
