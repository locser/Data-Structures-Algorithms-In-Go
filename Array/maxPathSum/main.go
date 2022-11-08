package main

import (
	"fmt"
	"time"
)

/*

Question#
Given two arrays in increasing order, find the maximum sum by choosing a few consecutive elements from one array and then a few elements from another array.
We can switch arrays at transition points, when the element value is the same in both arrays.

Input#
arr1 = [12, 13, 18, 20, 22, 26, 70]

arr2 = [11, 15, 18, 19, 20, 26, 30, 31]
Output#
//Max Sum elements: 11, 15, 18, 19, 20, 22, 26, 70

Max Sum: 201

*/
// my solution

func maxPathSum(arr1 []int, size1 int, arr2 []int, size2 int) int {

	max := 0
	var max1, max2, from1, to1, from2, to2 int
	for i := 0; i < size1; i++ {
		to1 = i
		max1 += arr1[i]
		index2 := contains(arr2, arr1[i])

		if index2 != -1 {
			fmt.Println("index:", index2)
			to2 = index2
			max1 = sumArr(arr1, from1, to1)
			fmt.Println(from1, to1, arr1[from1:to1+1], "-max1: ", max1)
			from1 = i + 1

			max2 = sumArr(arr2, from2, to2)
			fmt.Println(from2, to2, arr2[from2:to2+1], "-max2: ", max2)
			from2 = index2 + 1

			if max1 > max2 {
				max += max1
			} else {
				max += max2
			}
			fmt.Println("max:", max)
		}
	}
	max1 = sumArr(arr1, from1, size1-1)
	fmt.Println("Out For loop", from1, size1-1, arr1[from1:size1], "-max1: ", max1)

	max2 = sumArr(arr2, from2, size2-1)
	fmt.Println("Out For loop", from2, size2-1, arr2[from2:size2], "-max2: ", max2)

	if max1 > max2 {
		max = max + max1
	} else {
		max = max + max2
	}
	fmt.Println("max:", max)
	return max
}

func contains(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func sumArr(arr1 []int, from int, to int) int {
	max := 0

	for i := from; i <= to; i++ {
		max += arr1[i]
	}

	return max
}

// edu 's solution
/*
	No
*/

func main() {
	start := time.Now()
	array1 := []int{12, 13, 18, 20, 22, 26, 70}
	size1 := len(array1)

	array2 := []int{11, 15, 18, 19, 20, 26, 30, 31}
	size2 := len(array2)
	fmt.Println(maxPathSum(array1, size1, array2, size2))
	// elapsed := time.Since(start)  // you can count execution time in Go
	fmt.Println(time.Since(start))
}
