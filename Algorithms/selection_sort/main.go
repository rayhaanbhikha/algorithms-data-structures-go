package main

import "fmt"

func main() {
	numbers := []int{14, 33, 27, 10, 35, 19, 42, 44}
	sortedNums := selectionSort(numbers, 0)
	fmt.Println(sortedNums)
}

func indexOfSmallestNumInList(numbers []int, index int) int {
	smallestNumberIndex := index
	for i := index; i < len(numbers); i++ {
		if numbers[i] < numbers[smallestNumberIndex] {
			smallestNumberIndex = i
		}
	}
	return smallestNumberIndex
}

func selectionSort(numbers []int, index int) []int {
	for i := index; i < len(numbers); i++ {
		fmt.Println(numbers)
		if smallestNumIndex := indexOfSmallestNumInList(numbers, i); numbers[i] > numbers[smallestNumIndex] {
			fmt.Println("smallest num: ", smallestNumIndex)
			numbers[i], numbers[smallestNumIndex] = numbers[smallestNumIndex], numbers[i]
		}
	}
	return numbers
}
