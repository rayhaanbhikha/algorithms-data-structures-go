package main

import "fmt"

func main() {

	numbers := []int{9, 7, 8, 5, 3, 2, 1}
	sortedNums := mergeSort(numbers)
	fmt.Println(sortedNums)
}

func mergeSort(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	}
	mid := (len(numbers) - 1) / 2
	return merge(
		mergeSort(numbers[:mid+1]),
		mergeSort(numbers[mid+1:]),
	)
}

// [1] [2,7]
func merge(left []int, right []int) []int {
	fmt.Println("left: ", left, "  Right: ", right)
	result := make([]int, 0)
	leftIndex := 0
	rightIndex := 0
	lenOfLeft := len(left)
	lenOfRight := len(right)

	for leftIndex != lenOfLeft && rightIndex != lenOfRight {
		itemFromLeft := left[leftIndex]
		itemFromRight := right[rightIndex]
		if itemFromLeft < itemFromRight {
			result = append(result, itemFromLeft)
			leftIndex++
		} else {
			result = append(result, itemFromRight)
			rightIndex++
		}

		if leftIndex == lenOfLeft {
			result = append(result, right[rightIndex:]...)
			break
		}
		if rightIndex == lenOfRight {
			result = append(result, left[leftIndex:]...)
			break
		}
	}
	return result
}
