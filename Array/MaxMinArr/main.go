package main

import (
	"fmt"
	"time"
)

/*
Problem#
Given a sorted array, rearrange it into the maximum-minimum form.

Input#
A sorted integer array.

Output#
An integer array in a maximum-minimum form.

Sample input#
array = { 1, 2, 3, 4, 5, 6, 7 }
Sample output#
array = [ 7, 1, 6, 2, 5, 3, 4 ]

1, 		2, 		3, 		4, 		 5, 	 6, 	 7

7, 		1, 		6, 		2,		 5,		 3,		 4
max				max				max				max
		min				min				min
*/
//my solution
/*
index:
	0->1
	1->3
	2->5
	3->6
	4->4
	5->2
	6->1
*/
func MaxMinArr(arr []int, size int) {
	//Implement your solution here
	arr2 := make([]int, size)
	copy(arr2, arr)
	mid := size / 2
	for i := 0; i < size; i++ {
		if i < mid {
			arr[i*2+1] = arr2[i]
		} else { //if i >=mid
			arr[(size-i-1)*2] = arr2[i]
		}
	}
	fmt.Println(arr)

	//Note: Please update the same given array
	fmt.Println(arr2)
}

//edu
/*

func MaxMinArr(arr []int, size int) {
  aux := make([]int, size)
  copy(aux, arr)
  start := 0
  stop := size - 1
    for i := 0; i < size; i++ {
        if i%2 == 0 {
            arr[i] = aux[stop]
            stop -= 1
        } else {
            arr[i] = aux[start]
            start += 1
        }
    }
}


func swap(x, y int) (int, int) {
    return y, x
}

func MaxMinArr2(arr []int, size int) {
  for i := 0; i < (size - 1); i++ {
        ReverseArr(arr, i, size-1)
    }
}

func ReverseArr(arr []int, start int, stop int) {
  for start < stop {
        arr[start],arr[stop]=swap(arr[start], arr[stop])
        start += 1
        stop -= 1
  }
}

*/

func main() {
	start := time.Now()
	array := []int{1, 2, 3, 4, 5, 6, 7} // -> 7, 1, 6, 2,5,	3,4
	size := len(array)
	MaxMinArr(array, size)
	// elapsed := time.Since(start)  // you can count execution time in Go
	fmt.Println(time.Since(start))
}
