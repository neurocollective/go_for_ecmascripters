package main

import "fmt"

func main() {

	array := []int{ 1, 5, 16, 78, 505, 7890, 9986 }

	slice := make([]int, 7)

	fmt.Println("empty slice after creation:", slice)
	fmt.Println("slice from array[5:]", array[5:])
}
