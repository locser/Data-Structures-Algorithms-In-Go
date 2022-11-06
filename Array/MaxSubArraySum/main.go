package main

import "fmt"

/*
Problem#
Given an array of positive and negative integers, find a contiguous subarray whose sum (sum of its elements) is the maximum.

Input#
An integer array.

Output#
Maximum subarray sum.

Sample input#
array = { 1, -2, 3, 4, -4, 6, -14, 6, 2 }
Sample output#
sum = 9
Let’s look at the illustration to better understand how it works.

*/
func MaxSubArraySum(data []int) int {

	//Implement your solution here

	var max int
	high := len(data)

	for i := 0; i < high; i++ {
		maxtemp := data[i]
		for j := i + 1; j < high; j++ {
			maxtemp += data[j]
			if max <= maxtemp {
				max = maxtemp
			}
		}
	}
	return max
}

func MaxSubArraySum2(data []int) int {

	//Implement your solution here

	var max int
	size := len(data)

	maxtemp := 0
	for i := 0; i < size; i++ {
		maxtemp = maxtemp + data[i] //adding the values till current value
		if maxtemp < 0 {            // reset index if its negative
			maxtemp = 0
		}
		if max <= maxtemp {
			max = maxtemp
		}

	}
	return max
}

// go run main.go

func main() {
	array := []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
	fmt.Println(MaxSubArraySum(array))
	fmt.Println(MaxSubArraySum2(array))

}
