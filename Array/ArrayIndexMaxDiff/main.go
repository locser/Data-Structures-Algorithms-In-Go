package main

import (
	"fmt"
	"time"
)

/*
Problem#
Given an array arr[], find the pair of elements arr[i] and arr[j] such that the distance between i and j is maximum and arr[j] > arr[i].

Input#
An integer array.

Output#
Maximum distance.

Sample input#
array = { 33, 9, 10, 3, 2, 60, 30, 33, 1 }
Sample output#
ArrayIndexMaxDiff :  6

*/
// my solution

/*
	Sorry!
*/

//edu
/*
func ArrayIndexMaxDiff(arr []int, size int) int {
  maxDiff := -1
    j := 0
    for i := 0; i < size; i++ {
        j = size - 1
        for j > i {
            if arr[j] > arr[i] {
                if maxDiff < (j - i) {
                    maxDiff = (j - i)
                }

                break
              }
            j -= 1
          }

    }
    return maxDiff
}

*/
// edu 's solution
func ArrayIndexMaxDiff2(arr []int, size int) int {
	leftMin := make([]int, size)
	rightMax := make([]int, size)
	leftMin[0] = arr[0]
	i, j := 0, 0
	var maxDiff = 0
	for i = 1; i < size; i++ {
		if leftMin[i-1] < arr[i] {
			leftMin[i] = leftMin[i-1]
		} else {
			leftMin[i] = arr[i]
		}
	}
	rightMax[size-1] = arr[size-1]
	for i = size - 2; i >= 0; i-- {
		if rightMax[i+1] > arr[i] {
			rightMax[i] = rightMax[i+1]
		} else {
			rightMax[i] = arr[i]
		}
	}
	i = 0
	j = 0
	maxDiff = -1
	for j < size && i < size {
		if leftMin[i] < rightMax[j] {
			if maxDiff < j-i {
				maxDiff = j - i
			}
			j = j + 1
		} else {
			i = i + 1
		}
	}
	return maxDiff
}

func main() {
	start := time.Now()
	array := []int{1, 2, 3, 4, 5, 6, 7} // -> 7, 1, 6, 2,5,	3,4
	size := len(array)
	fmt.Println("ArrayIndexMaxDiff : ", ArrayIndexMaxDiff2(array, size))
	// elapsed := time.Since(start)  // you can count execution time in Go
	fmt.Println(time.Since(start))
}
