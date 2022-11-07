package main

import (
	"fmt"
	"sort"
	"time"
)

/*
PProblem#
Given an unsorted array, find the smallest positive number missing in the array.

Input#
An integer array.

Output#
A missing integer. Otherwise -1 if not found.

Sample input#
array = { 8, 5, 6, 1, 9, 11, 2, 7, 4, 10 }
Sample output#
missing number = 3
*/
//my solution
// import "sort"
func SmallestPositiveMissingNumber(arr []int, size int) int {
	//Implement your solution here
	sort.Ints(arr[:])
	for i := 0; i < size; i++ {
		if i+1 != arr[i] {
			return i + 1
		}
	}
	return -1 //Return -1 if missing number not found
}

//edu
/*
func SmallestPositiveMissingNumber3(arr []int, size int) int {
     aux := make([]int, size)
      for index, _ := range aux {
        aux[index] = -1
      }
    for i := 0; i < size; i++ {
      if arr[i] > 0 && arr[i] <= size {
          aux[arr[i]-1] = arr[i]
      }
    }
    for i := 0; i < size; i++ {
        if aux[i] != i+1 {
          return i + 1
        }
    }
return -1
}

func SmallestPositiveMissingNumber2(arr []int, size int) int {
  hs := make(map[int]int)
  for i := 0; i < size; i++ {
        hs[arr[i]] = 1
  }
    for i := 1; i < size+1; i++ {
        _, ok := hs[i]
        if ok == false {
            return i
        }
    }

return -1
}
*/

// go run main.go

func main() {
	start := time.Now()
	array := []int{5, 2, 1, 4}
	size := len(array)
	fmt.Println(SmallestPositiveMissingNumber(array, size))
	elapsed := time.Since(start) // you can count execution time in Go
	fmt.Println(elapsed)
}
