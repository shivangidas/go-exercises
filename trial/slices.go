package main

import (
	"fmt"
	"sort"
)

func addNum(nums []int, newNum int) {
	nums = append(nums, newNum)
	fmt.Printf("addNum nums %v\n", nums)
}

func sortNum(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Printf("sortNum nums %v\n", nums)
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Printf("  main nums %v\n", nums)
	addNum(nums, 4)
	fmt.Printf("  main nums %v\n", nums)
	sortNum(nums)
	fmt.Printf("  main nums %v\n", nums)
}
