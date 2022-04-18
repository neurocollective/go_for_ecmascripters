package main

import "fmt"

func main() {

	mappy := map[string]string {
		"key": "value",
		"dude": "brah",
	}

	fmt.Println("checking value of \"brah\" => ")
	// go's multiple return values
	value, present := mappy["dude"]

	if present {
		fmt.Println("value is", value)
	} else {
		fmt.Println("not in map!")
	}
}
