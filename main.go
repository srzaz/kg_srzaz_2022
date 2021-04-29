package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]
	numbers := [10]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	arg := os.Args[3]

	fmt.Println(numbers)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
