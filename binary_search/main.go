package main

import "fmt"

func main() {
	items := []int{1, 3, 5, 7, 9}
	// fmt.Print(5 / 2)
	fmt.Println(binarySearchR(items, -1))
	// fmt.Println(*binarySearch(items, 3))
	// fmt.Println(binarySearch(items, -1))
}

func binarySearch(items []int, item int) *int {
	n := len(items)
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		guess := items[mid]
		switch {
		case guess == item:
			return &mid
		case guess < item:
			low = mid + 1
		case guess > item:
			high = mid - 1
		}
	}
	return nil
}

func binarySearchR(items []int, item int) int {
	// base case
	if len(items) == 1 {
		if item == items[0] {
			return 1
		}
		panic("can't find it")
	}
	// recursive case
	low := 0
	high := len(items) - 1
	mid := (low + high)/2
	guess := items[mid]
	newItemList := make([]int, 0)
	if guess == item {
		return item
	} else if guess > item {
		newItemList = items[:mid]
	} else {
		newItemList = items[mid+1:]
	}
	return binarySearchR(newItemList, item)
}