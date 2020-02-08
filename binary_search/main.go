package main

import "fmt"

func main() {
	items := []int{1, 3, 5, 7, 9}
	// fmt.Print(5 / 2)
	fmt.Println(*binarySearch(items, 3))
	fmt.Println(binarySearch(items, -1))
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
