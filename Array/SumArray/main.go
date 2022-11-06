package main

func sumArr(data []int) int {

	total := 0

	//Implement your solution here
	for _, v := range data {
		total += v
	}

	return total
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	println(sumArr(array))

}
