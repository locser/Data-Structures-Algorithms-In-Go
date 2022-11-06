package main

import "fmt"

/*
Problem#
Given an data, rotate its elements k times.

Input#
An integer data and an integer k.

Output#
A rotated data.

Sample input#
data = { 1,2,3,4,5,6 }
k = 2
Sample output#
data = { 3,4,5,6,1,2 }

*/
//my solution
func RotateArray2(data []int, k int) {
	//Implement your solution here
	size := len(data)
	for i := 0; i < k; i++ {
		first := data[0]

		copy(data, data[1:])
		data[size-1] = first
	}

	//You must make changes in the given array

	fmt.Println(data)
}

//edu
func RotateArray(data []int, k int) {
	n := len(data)
	// first part
	ReverseArray(data, 0, k-1) // reverse 1st part of array till kth position
	fmt.Println(data)          //opt
	ReverseArray(data, k, n-1) // reverse from kth position till end
	fmt.Println(data)          //opt
	//second part
	ReverseArray(data, 0, n-1) // reverse whole array
}

func ReverseArray(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

// go run main.go

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	RotateArray(data, 2)

}
