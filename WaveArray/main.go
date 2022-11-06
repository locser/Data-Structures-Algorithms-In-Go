package main

import (
	"fmt"
	"sort"
)

/*
Problem#
Given an array of integers, arrange the array elements in a wavefrom such that the element at the odd indices are less than or equal to their neighboring elements at the even indices.

Input#
An integer array.

Output#
An array in a waveform.

Sample input#
array = { 8, 1, 2, 3, 4, 5, 6, 4, 2 }

Sample output#
array = [ 8, 1, 3, 2, 5, 4, 6, 2, 4 ]
OR
array = [ 2, 1, 3, 2, 4, 4, 6, 5, 8 ]

*/
//my solution
func WaveArray(arr []int) {
	//Implement your solution here
	sort.Ints(arr[:])
	//Update the same array
	size := len(arr)
	for i := 0; i < size-2; i = i + 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
	fmt.Println(arr)
}

//edu

// go run main.go

func main() {
	array := []int{8, 1, 2, 3, 4, 5, 6, 4, 2}
	WaveArray(array)
}
