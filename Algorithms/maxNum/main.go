package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 10, 3}
	numbers := []int{1, 2, 3, 10, 54, 36, 8, 15, 7, 30, 5}
	fmt.Println(maxNum(numbers))
}

func maxNum(numbers []int) int {
	// base case
	if len(numbers) == 2 {
		if numbers[0] > numbers[1] {
			return numbers[0]
		}
		return numbers[1]
	}

	// recursive case
	subMax := maxNum(numbers[1:])
	if subMax > numbers[0] {
		return subMax
	}
	return numbers[0]
}
