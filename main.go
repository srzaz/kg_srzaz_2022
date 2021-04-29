package main

import (
	"fmt"
	"os"
	"strconv"
)

//Program for converting integer array to
//array of the english form of the integer

func main() {

	input := os.Args[1:]

	//error handling for input
	var err error
	eng_numbers := [10]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	nums := make([]int, len(input))

	//convert command line arguments into integer form
	for i := 0; i < len(input); i++ {
		if nums[i], err = strconv.Atoi(input[i]); err != nil {
			panic(err)
		}
	}
	//loop through numbers and determine the digit through modulo and division
	//takes 0 <= n < 1,000,000
	for i := 0; i < len(nums); i++ {
		engnum := ""
		hunthous := eng_numbers[nums[i]/100000]
		tenthous := eng_numbers[(nums[i]%100000)/10000]
		thous := eng_numbers[(nums[i]%10000)/1000]
		hun := eng_numbers[(nums[i]%1000)/100]
		ten := eng_numbers[(nums[i]%100)/10]
		one := eng_numbers[(nums[i] % 10)]

		if nums[i] >= 100000 {
			engnum += hunthous + tenthous + thous + hun + ten + one
		} else if nums[i] >= 10000 {
			engnum += thous + hun + ten + one
		} else if nums[i] >= 1000 {
			engnum += thous + hun + ten + one
		} else if nums[i] >= 100 {
			engnum += hun + ten + one
		} else if nums[i] >= 10 {
			engnum += ten + one
		} else {
			engnum += one
		}
		if i != (len(nums) - 1) {
			fmt.Print(engnum + ", ")
		} else {
			fmt.Println(engnum)
		}

	}

}
