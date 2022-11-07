package main

/*
Problem#
You are given an array of size n, containing elements from indexes 0 to n-1. All values from 0 to n-1 are present in the array, and if a value is not in the array, then -1 is there to take its place. Arrange values of the array such that value i is stored at array[i].

Input#
An integer array of numbers ranging from 0 to n-1 with -1 for numbers missing in this range.

Constraints

No element in the input array is greater than n-1.
The elements given in the input array are unique except -1.
Output#
A sorted integer array.

Sample input#
array = { 8, -1, 6, 1, 9, 3, 2, 7, 4, -1 }
Sample output#
array = [ -1, 1, 2, 3, 4, -1, 6, 7, 8, 9 ]
*/
//my solution
//no

//edu

// go run main.go

import (
	"fmt"
)

func indexArray(arr []int, size int) {
	for i := 0; i < size; i++ {
		curr := i
		value := -1
		fmt.Println(curr, value)
		/* swaps to move elements in the proper position. */
		for arr[curr] != -1 && arr[curr] != curr {

			temp := arr[curr]
			arr[curr] = value
			value = temp
			curr = temp
			fmt.Println(curr, arr[curr], arr)
		}
		/* check if some swaps happened. */
		if value != -1 {
			arr[curr] = value
		}
	}
}

// second solution
// func indexArray2(arr []int, size int) {
// 	temp := 0
// 	for i := 0; i < size; i++ {
// 	   for arr[i] != -1 && arr[i] != i {
// 	   /* swap arr[i] and arr[arr[i]] */
// 		   temp = arr[i]
// 		   arr[i] = arr[temp]
// 		   arr[temp] = temp
// 	   }
//    }
// }

/* Testing code */
func main() {
	arr := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
	size := len(arr)
	indexArray(arr, size)
	fmt.Println(arr)
}
