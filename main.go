package main

import "fmt"

func main() {
	numbers := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	sortedNums := bubbleSort(numbers)
	fmt.Println(sortedNums)
}

func bubbleSort(numbers []int) []int {
	swapped := false

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			swapped = true
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
		}
	}
	fmt.Println(numbers)
	if swapped {
		return bubbleSort(numbers)
	}

	return numbers
}
