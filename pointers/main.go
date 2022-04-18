package main

import "fmt"


func mutateFail(person Person) {
	person.FirstName = "Bruh"
}

func mutate(address *Person) {
	(*address).FirstName = "Bruh"
}

type Person struct {
	FirstName string
	LastName string
}

func main() {

	var text string = "sup brah"
	var pointer *string = &text

	fmt.Println("text", text)
	fmt.Println("pointer", pointer)

	// text = nil
	pointer = nil

	// fmt.Println("text", text)
	fmt.Println("pointer", pointer)

	/* mutation */

	dude := Person { "Cody", "McCoder"}

	mutateFail(dude)
	fmt.Println("after mutateFail:", dude)

	mutate(&dude)
	fmt.Println("after mutate:", dude)
}
