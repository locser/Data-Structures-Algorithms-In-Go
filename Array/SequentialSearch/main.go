package main

func SequentialSearch(data []int, value int) bool {
	//Implement your solution here
	for _, v := range data {
		if v == value {
			return true
		}
	}
	return false //Return false if value not found
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	key := 3
	println(SequentialSearch(array, key))
	// go run main.go
}
