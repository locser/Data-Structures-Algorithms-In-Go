package main

import (
	"fmt"
	"time"
)

/*
Problem#
For a given value n find n!, where n! = n∗(n−1)⋯2∗1
. Solve this problem using recursion.

Input#
A positive integer.

Output#
A factorial.

Sample input#
5
Sample output#
Answer: 120 // 5*4*3*2*1
Coding exercise#
Try to solve this yourself first. If you get stuck or need help, you can always press the “Show Solution” button to see how your problem can be solved. We’ll look at the solution in the next lesson.

Good luck!

*/

func Factorial(i int) int {
	// Implement your solution here
	if i == 1 {
		return 1
	}
	return i * Factorial(i-1)
}

func main() {
	start := time.Now()
	fmt.Println("Factorial : ", Factorial(5))
	// elapsed := time.Since(start)  // you can count execution time in Go
	fmt.Println(time.Since(start))
}
