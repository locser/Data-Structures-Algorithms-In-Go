package main

import (
	"fmt"
	"time"
)

/*
Problem#
Given an integer in decimal form, print its hexadecimal form. Use recursion to solve the problem.

Input#
An integer.

Output#
A string of the hexadecimal form of an integer.

Sample input#
95
Sample output#
5F

*/

/* Explanation#
Line 36: We initialize a conversion variable that we’ll use to convert a base-10 digit to base-16.
Line 37: We define a base variable to use later in the code.
Line 38: We take the modulus of number using the base variable to separate a base-16 digit from the number and save it to the digit variable.
Line 39: We divide the number with base to make changes in the number variable too.
Line 40–42: We first check if number is not equal to 0. If true, the printInt(number) is called recursively. If false, the function will not be called recursively and works as a base case.
Line 43: After exiting the if condition, the digit variable is used as an index of the conversion variable to finally convert the base 10 number to the base 16 number. Finally, we print the converted digit to the console.

*/

func printInt(number int) {
	conversion := "0123456789ABCDEF"
	base := 16
	digit := number % base
	number = number / base
	if number != 0 {
		printInt(number)
	}
	fmt.Println(string(conversion[digit]))
}

func main() {
	start := time.Now()
	printInt(95)
	// elapsed := time.Since(start)  // you can count execution time in Go
	fmt.Println(time.Since(start))
}
