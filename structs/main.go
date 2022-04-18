package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
}

func main() {

	dude := Person { "Cody", "McCoder"}

	fmt.Println("Hello! My name is", dude.FirstName, dude.LastName)

}
