package main

import "fmt"

func main() {

	ingredients := []string{ "bacon bits", "ranch dressing", "no lettuce" }

	fmt.Println("Here's how to make the universe's greatest salad:")

	for index, value := range ingredients {
		fmt.Println("- step", index + 1, "add", value)
	}

	fmt.Println("And that's it! Vegetables are the worst LOLZ!")
	fmt.Println("(dies from malnutrition)")
}
