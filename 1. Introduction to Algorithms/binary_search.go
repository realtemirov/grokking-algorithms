package main

import (
	"fmt"
)

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	item := 9
	fmt.Println(binary_search(list, item))
}

func binary_search(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {

		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
