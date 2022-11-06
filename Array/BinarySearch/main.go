package main

import "fmt"

/*
Explanation#
The binary search algorithm is used to find a specific value in the sorted list. At each step, we look at the middle index. If the item at the middle index is the same as the desired one, it is returned. Otherwise, the middle value is larger or smaller than the value we’re looking for. If our desired value is smaller, we confine our search space to the left half of the array and ignore the right half and vice versa. We prune half of the search space at each stage. Because of that,this algorithm is more efficient than the linear search algorithm.

Problem#
Find a value in a sorted array using the binary search algorithm.

Input#
An integer array and a key.

Output#
Return true if the key is found, false otherwise.

Sample input#
array = { 1, 2, 3, 4, 5, 6, 7, 8 }
key = 3
Sample output#
true

*/
func BinarySearch(data []int, value int) bool {
	//Implement your solution here
	var mid int
	low := 0
	high := len(data)
	for low <= high {
		mid = (high + low) / 2
		if data[mid] == value {
			return true
		} else {
			if data[mid] > value {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

	}
	return false
}

// go run main.go

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	key := 3
	fmt.Println(BinarySearch(array, key))
	// println(BinarySearch(array, key))
}
