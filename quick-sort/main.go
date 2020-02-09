package main

import "fmt"

func main() {
	numbers := []int{9, 7, 8, 5, 3, 2, 1}
	fmt.Println(quickSort(numbers))
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[(len(nums)-1)/2]
	subArray := make([]int, 0)
	supArray := make([]int, 0)

	for _, num := range nums {
		if num != pivot {
			if num < pivot {
				subArray = append(subArray, num)
			} else {
				supArray = append(supArray, num)
			}
		}
	}

	sortedArray := quickSort(subArray)
	sortedArray = append(sortedArray, pivot)
	sortedArray = append(sortedArray, quickSort(supArray)...)

	return sortedArray

}
