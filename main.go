package main

import (
	"fmt"
	"os"
	"strconv"
)

//Program for converting integer array to
//array of the english form of the integer
//determine the digit through modulo and division
func intToString(num int) string {
	eng_numbers := [10]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	engnum := ""
	hunthous := eng_numbers[num/100000]
	tenthous := eng_numbers[(num%100000)/10000]
	thous := eng_numbers[(num%10000)/1000]
	hun := eng_numbers[(num%1000)/100]
	ten := eng_numbers[(num%100)/10]
	one := eng_numbers[(num % 10)]

	//determine size of integer
	if num >= 100000 {
		engnum += hunthous + tenthous + thous + hun + ten + one
	} else if num >= 10000 {
		engnum += thous + hun + ten + one
	} else if num >= 1000 {
		engnum += thous + hun + ten + one
	} else if num >= 100 {
		engnum += hun + ten + one
	} else if num >= 10 {
		engnum += ten + one
	} else {
		engnum += one
	}
	return engnum
}

func main() {

	input := os.Args[1:]

	//error handling for input
	var err error

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
		result := intToString(nums[i])

		if i != (len(nums) - 1) {
			fmt.Print(result + ", ")
		} else {
			fmt.Println(result)
		}
	}

}
