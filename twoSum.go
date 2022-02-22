package main

import "fmt"

var (
	nums = []int{1, 2, 3}
)

func main() {
	twoSum(nums, 3)
}

/*This will traverse a slice of ints, hash the values of a list[nums] in a map,
and return the indicies of the values that add to the target number */
func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, v := range nums {
		_, ok := hm[v]
		if ok {
			idx := []int{i, hm[v]}
			fmt.Println(idx)
			return idx
		}
		hm[target-v] = i
	}

	return nil
}
