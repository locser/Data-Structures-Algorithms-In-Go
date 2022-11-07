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

	return -1
}

//edu
/*


 */
// edu 's solution

func main() {
	start := time.Now()
	array1 := []int{1, 2, 3, 4, 5, 6, 7} // -> 7, 1, 6, 2,5,	3,4
	size1 := len(array1)

	array2 := []int{1, 2, 3, 4, 5, 6, 7} // -> 7, 1, 6, 2,5,	3,4
	size2 := len(array2)
	_, _, _, _ = array1, array2, size1, size2
	fmt.Println(maxPathSum(array1, size1, array2, size2))
	// elapsed := time.Since(start)  // you can count execution time in Go
	fmt.Println(time.Since(start))
}
