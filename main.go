package main

import "fmt"

func value(r string) int {

	if r == "I" {
		return 1
	}

	if r == "V" {
		return 5
	}

	if r == "X" {
		return 10
	}

	if r == "L" {
		return 50
	}

	if r == "C" {
		return 100
	}

	if r == "D" {
		return 500
	}

	if r == "M" {
		return 1000
	} else {

		return -1
	}

}

func romantonumber(roman string) int {

	// Initialize result
	res := 0

	// Traverse given input
	for i := 0; i < len(roman); i++ {
		// Getting value of symbol s[i]
		s1 := value(string(roman[i]))

		if i+1 < len(roman) {
			// Getting value of symbol s[i+1]
			s2 := value(string(roman[i+1]))

			// Comparing both values
			if s1 >= s2 {
				// Value of current symbol is greater
				// or equal to the next symbol
				res = res + s1
			} else {
				res = res + s2 - s1
				i++ // Value of current symbol is
				// less than the next symbol
			}
		} else {
			res = res + s1

		}
	}
	return res
}

func main() {
	var str string
	fmt.Println("enter a roman numeral value")
	fmt.Scanln(&str)

	//str = "IV"
	fmt.Println("Integer form of Roman Numeral is ", romantonumber(str))

}
