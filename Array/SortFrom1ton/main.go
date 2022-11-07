package main

import (
	"fmt"
	"sort"
	"time"
)

/*
Problem#
You are given an array of length n
n
. It contains unique elements from 1 to n
n
. You are required to sort the array elements. Missing elements will be placed at -1.

Input#
An integer array.

Output#
A sorted array.

Sample input#
array = { 8, 5, 6, 1, 9, 3, 2, 7, 4, 10 }
Sample output#
array = [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ]

*/
//my solution
// import "sort"
func Sort1toN(arr []int, size int) {

	//Implement your solution here
	sort.Ints(arr[:])
	fmt.Println(arr)

}

//edu
/*
func Sort1toN1(arr []int, size int) {
    curr, value, next := 0, 0, 0
    for i := 0; i < size; i++ {
      curr = i
      value = -1
      swaps to move elements in the proper position.
      for curr >= 0 && curr < size && arr[curr] != curr+1 {
        next = arr[curr]
        arr[curr] = value
        value = next
        curr = next - 1

      }
    }
}
func Sort1toN2(arr []int, size int) {
    temp := 0
    for i := 0; i < size; i++ {
        for arr[i] != i+1 && arr[i] > 1 {
            // simple swapping with correct index
            temp = arr[i]
            arr[i] = arr[temp-1]
            arr[temp-1] = temp
        }
 }
}
*/

// go run main.go

func main() {
	start := time.Now()
	array := []int{5, 3, 2, 1, 4}
	size := len(array)
	Sort1toN(array, size)
	elapsed := time.Since(start) // you can count execution time in Go
	fmt.Println(elapsed)
}
