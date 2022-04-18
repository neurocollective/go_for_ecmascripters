package main

import (
	"fmt"
	"os"
)


func main() {
    data, err := os.ReadFile("./bruh.txt")
    if err != nil {
    	fmt.Println("oh noes!", err.Error())
    	return
    }

    fmt.Println("bruh.txt contents =>")
    fmt.Println(string(data))
}