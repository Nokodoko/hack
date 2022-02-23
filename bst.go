package main

import "fmt"

var (
	list = []int{1, 2, 3, 4, 5}
)

func search(list []int, target int) int {
	low, high := 0, len(list)-1
	for low <= high {
		mid := low + (high-low)>>1
		switch {
		case list[mid] == target:
			fmt.Println(mid)
			return mid
		case list[mid] > target:
			high = mid - 1
		default:
			low = mid + 1
		}
	}
	fmt.Println(-1)
	return -1
}
